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
)

var svnlogRegex *regexp.Regexp

type SvnLog struct {
	Reversion string
	Author    string
	Time      string
	Content   string
}

type SvnLogList []SvnLog

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
	ret := deply(r.PostFormValue("groupid"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", strings.Replace(ret, "{CRLF}", "\n", -1)))
}

func Rollback(w http.ResponseWriter, r *http.Request) {
	ret := rollback(r.PostFormValue("groupid"), r.PostFormValue("reversion"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", strings.Replace(ret, "{CRLF}", "\n", -1)))
}

func ShowLog(w http.ResponseWriter, r *http.Request) {
	list, err := showlog(r.PostFormValue("groupid"))
	if err != nil {
		fmt.Fprintf(w, "%s\n", helper.JsonResp(false, err.Error(), nil))
	} else {
		fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", list))
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
					go sendAndExecJob(jobBody, jobExecChan, node.Addr, app.Name, node.Alias)
				}
			}

			resp := ""
			for i := 0; i < chanLen; i++ {
				exeRet := <-jobExecChan
				if exeRet.Err != nil {
					resp += fmt.Sprintf("[%s]ERROR => %s", exeRet.NodeName, exeRet.Err.Error())
				} else {
					resp += fmt.Sprintf("[%s]%s", exeRet.NodeName, exeRet.Message)
				}
			}
			return resp
		}
	}
	return ""
}

//send job to client and get execute result
func sendAndExecJob(jobBody []byte, jobExecChan chan jobExecResult, addr string, appName string, nodeName string) {
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
					go sendAndExecJob(jobBody, jobExecChan, node.Addr, app.Name, node.Alias)
				}
			}

			resp := ""
			for i := 0; i < chanLen; i++ {
				exeRet := <-jobExecChan
				if exeRet.Err != nil {
					resp += fmt.Sprintf("[%s]ERROR => %s", exeRet.NodeName, exeRet.Err.Error())
				} else {
					resp += fmt.Sprintf("[%s]%s", exeRet.NodeName, exeRet.Message)
				}
			}
			return resp
		}
	}
	return ""
}

//read svn log
func showlog(groupid string) (list SvnLogList, err error) {
	groupid = strings.TrimSpace(groupid)
	if groupid != "" {
		for _, app := range config.C.Apps {
			if app.GroupId == groupid {
				bytes, err := helper.RunShell(fmt.Sprintf("svn log --limit 100 %s", app.Url))
				if err != nil {
					return nil, err
				} else {
					match := svnlogRegex.FindAllSubmatch(bytes, -1)
					logList := make(SvnLogList, 0)
					for _, item := range match {
						svnlog := SvnLog{Reversion: string(item[1]), Author: string(item[2]), Time: string(item[3]), Content: string(item[4])}
						logList = append(logList, svnlog)
					}
					return logList, nil
				}
			}
		}
	}
	return nil, errors.New("groupid invalid")
}
