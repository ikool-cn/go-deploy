package ctrl

import (
	"net/http"
	"fmt"
	"go-deploy/tpl"
	"go-deploy/config"
	"time"
	"log"
	"encoding/json"
	"go-deploy/helper"
	"strings"
	"errors"
	"regexp"
	"net"
	"bufio"
	"strconv"
)

var svnlogRegex *regexp.Regexp
var separator = "{CRLF}"
var dataFormat = "2006-01-02 15:04:05"

type LogEntity struct {
	Reversion string
	Author    string
	Time      string
	Content   string
}

type LogList []LogEntity

type jobExecResult struct {
	Err      error
	Message  string
	AppName  string
	NodeName string
}

func init() {
	svnlogRegex = regexp.MustCompile(`r(\d+) \| (\w+) \| (.*) \+0800(?:.*)\n\n(.*)\n--`)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", tpl.GetIndexTpl())
}

func List(w http.ResponseWriter, r *http.Request) {
	str, _ := json.Marshal(config.C.Apps)
	fmt.Fprintf(w, "%s\n", str)
}

func Deply(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ret := deply(r.PostFormValue("groupid"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", strconv.FormatFloat(time.Since(start).Seconds(), 'f', 3, 64), strings.Replace(ret, separator, "\n", -1)))
}

func Rollback(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ret := rollback(r.PostFormValue("groupid"), r.PostFormValue("reversion"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", strconv.FormatFloat(time.Since(start).Seconds(), 'f', 3, 64), strings.Replace(ret, separator, "\n", -1)))
}

func ShowLog(w http.ResponseWriter, r *http.Request) {
	list, err := showlog(r.PostFormValue("groupid"))
	if err != nil {
		fmt.Fprintf(w, "%s\n", helper.JsonResp(false, err.Error(), "", nil))
	} else {
		fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", "", list))
	}
}

//send a update message to the group nodes
func deply(groupid string) string {
	for _, app := range config.C.Apps {
		if app.GroupId == groupid {
			jobExecChan := make(chan jobExecResult, len(app.Node))
			chanLen := 0

			for _, node := range app.Node {
				if node.Online {
					jobBody, _ := json.Marshal(struct {
						Type        string `json:"type"`
						Path        string `json:"path"`
						Action      string `json:"action"`
						BeforDeploy string `json:"befor_deploy"`
						AfterDeploy string `json:"after_deploy"`
					}{Type: node.Type, Path: node.Path, BeforDeploy: node.BeforDeploy, AfterDeploy: node.AfterDeploy, Action: "update"})

					chanLen ++
					go dispatchJob(jobBody, jobExecChan, node.Addr, app.Name, node.Alias)
				}
			}

			resp := ""
			for i := 0; i < chanLen; i++ {
				exeRet := <-jobExecChan
				if exeRet.Err != nil {
					resp += fmt.Sprintf("[%s:%s]ERROR => %s", exeRet.AppName, exeRet.NodeName, exeRet.Err.Error())
				} else {
					resp += fmt.Sprintf("[%s:%s]%s", exeRet.AppName, exeRet.NodeName, exeRet.Message)
				}
			}
			return resp
		}
	}
	return ""
}

//send job to client and get execute result
func dispatchJob(jobBody []byte, jobExecChan chan jobExecResult, addr string, appName string, nodeName string) {
	execResult := jobExecResult{AppName: appName, NodeName: nodeName}
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("Error connect to client:", err)
		execResult.Err = err
		jobExecChan <- execResult
		return
	}
	defer conn.Close()

	jobBody = append(jobBody, '\n')
	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	_, err = conn.Write(jobBody)
	if err != nil {
		log.Println("Error writing to stream:", err)
		execResult.Err = err
		jobExecChan <- execResult
		return
	}

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Client closed", err.Error())
		execResult.Err = err
		jobExecChan <- execResult
		return
	}
	execResult.Message = message
	jobExecChan <- execResult
	return
}

//send a rollback message to the group nodes
func rollback(groupid string, reversion string) string {
	for _, app := range config.C.Apps {
		if app.GroupId == groupid {
			jobExecChan := make(chan jobExecResult, len(app.Node))
			chanLen := 0

			for _, node := range app.Node {
				if node.Online {
					jobBody, _ := json.Marshal(struct {
						Type        string `json:"type"`
						Path        string `json:"path"`
						Action      string `json:"action"`
						Reversion   string `json:"reversion"`
						BeforDeploy string `json:"befor_deploy"`
						AfterDeploy string `json:"after_deploy"`
					}{Type: node.Type, Path: node.Path, BeforDeploy: node.BeforDeploy, AfterDeploy: node.AfterDeploy, Action: "rollback", Reversion: reversion})

					chanLen ++
					go dispatchJob(jobBody, jobExecChan, node.Addr, app.Name, node.Alias)
				}
			}

			resp := ""
			for i := 0; i < chanLen; i++ {
				exeRet := <-jobExecChan
				if exeRet.Err != nil {
					resp += fmt.Sprintf("[%s:%s]ERROR => %s", exeRet.AppName, exeRet.NodeName, exeRet.Err.Error())
				} else {
					resp += fmt.Sprintf("[%s:%s]%s", exeRet.AppName, exeRet.NodeName, exeRet.Message)
				}
			}
			return resp
		}
	}
	return ""
}

func showlog(groupid string) (list LogList, err error) {
	groupid = strings.TrimSpace(groupid)
	if groupid != "" {
		for _, app := range config.C.Apps {
			if app.GroupId == groupid {
				if app.Type == "svn" {
					return showSvnLog(app)
				} else if app.Type == "git" {
					return showGitLog(app)
				}
			}
		}
	}
	return nil, errors.New("groupid invalid")
}

//svn log --limit 100 svn://x.x.x.x/path
func showSvnLog(app config.Apps) (list LogList, err error) {
	bytes, err := helper.RunShell(fmt.Sprintf("svn log --limit 100 %s", app.Url))
	if err != nil {
		return nil, err
	} else {
		match := svnlogRegex.FindAllSubmatch(bytes, -1)
		logList := make(LogList, 0)
		for _, item := range match {
			svnlog := LogEntity{Reversion: string(item[1]), Author: string(item[2]), Time: string(item[3]), Content: string(item[4])}
			logList = append(logList, svnlog)
		}
		return logList, nil
	}
}

//cd /pathto/xx && git log -50 --pretty="%h {CRLF} %an {CRLF} %at {CRLF} %s"
func showGitLog(app config.Apps) (list LogList, err error) {
	cmd := fmt.Sprintf(`cd %s && git log -100 --pretty="%%h %s %%an %s %%at %s %%s"`, app.Fetchlogpath, separator, separator, separator)
	bytes, err := helper.RunShell(cmd)
	if err != nil {
		return nil, err
	} else {
		logs := strings.Split(string(bytes), "\n")
		logList := make(LogList, 0)
		for _, line := range logs {
			if strings.TrimSpace(line) != "" {
				fmt.Println(line)
				commitLog := strings.Split(line, separator)
				timeInt64, err := strconv.ParseInt(strings.TrimSpace(commitLog[2]), 10, 64)
				if err != nil {
					timeInt64 = time.Now().Unix()
				}
				logList = append(logList, LogEntity{Reversion: commitLog[0], Author: commitLog[1], Time: time.Unix(timeInt64, 0).Format(dataFormat), Content: commitLog[3]})
			}
		}
		return logList, nil
	}
}
