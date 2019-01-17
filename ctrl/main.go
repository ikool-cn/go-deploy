package ctrl

import (
	"net/http"
	"fmt"
	"go-deploy/tpl"
	"go-deploy/config"
	"net"
	"time"
	"log"
	"encoding/json"
	"go-deploy/helper"
	"strings"
	"errors"
	"regexp"
)

var svnlogRegex *regexp.Regexp

type SvnLog struct {
	Reversion string
	Author    string
	Time      string
	Content   string
}

type SvnLogList []SvnLog

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
	deply(r.PostFormValue("groupid"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", nil))
}

func Rollback(w http.ResponseWriter, r *http.Request) {
	rollback(r.PostFormValue("groupid"), r.PostFormValue("reversion"))
	fmt.Fprintf(w, "%s\n", helper.JsonResp(true, "", nil))
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
func deply(groupid string) {
	for _, app := range config.C.Apps {
		if app.GroupId == groupid {
			for _, node := range app.Node {
				if node.Online && node.Conn != nil {
					bytes, _ := json.Marshal(struct {
						Type        string `json:"type"`
						Path        string `json:"path"`
						Action      string `json:"action"`
						BeforDeploy string `json:"befor_deploy"`
						AfterDeploy string `json:"after_deploy"`
					}{Type: node.Type, Path: node.Path, BeforDeploy: node.BeforDeploy, AfterDeploy: node.AfterDeploy, Action: "update"})
					go func(conn net.Conn, bytes []byte) {
						bytes = append(bytes, '\n')
						conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
						_, err := conn.Write(bytes)
						if err != nil {
							log.Println("deply err", err)
						}
					}(node.Conn, bytes)
				}
			}
			return
		}
	}
}

//send a rollback message to the group nodes
func rollback(groupid string, reversion string) {
	for _, app := range config.C.Apps {
		if app.GroupId == groupid {
			for _, node := range app.Node {
				if node.Online && node.Conn != nil {
					bytes, _ := json.Marshal(struct {
						Type        string `json:"type"`
						Path        string `json:"path"`
						Action      string `json:"action"`
						Reversion   string `json:"reversion"`
						BeforDeploy string `json:"befor_deploy"`
						AfterDeploy string `json:"after_deploy"`
					}{Type: node.Type, Path: node.Path, BeforDeploy: node.BeforDeploy, AfterDeploy: node.AfterDeploy, Action: "rollback", Reversion: reversion})
					go func(conn net.Conn, bytes []byte) {
						bytes = append(bytes, '\n')
						conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
						_, err := conn.Write(bytes)
						if err != nil {
							log.Println("deply err", err)
						}
					}(node.Conn, bytes)
				}
			}
			return
		}
	}
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
