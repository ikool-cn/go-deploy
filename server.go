package main

import (
	"net"
	"os"
	"bufio"
	"strings"
	"time"
	"log"
	"path/filepath"
	"deploy/config"
	"github.com/bwmarrin/snowflake"
	"net/http"
	"fmt"
	"deploy/tpl"
	"encoding/json"
)

var cfg *config.Config

func init() {
	root, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := filepath.Join(root, "config.json")
	cfg = config.New(file)
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	for key := range cfg.Apps {
		cfg.Apps[key].GroupId = node.Generate().String()
	}
}

func main() {
	//start http server
	log.Printf("Http server start,listening at %s", cfg.ListenHttp)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", tpl.GetIndexTpl())
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		str, _ := json.Marshal(cfg.Apps)
		fmt.Fprintf(w, "%s\n", str)
	})

	http.HandleFunc("/deply", func(w http.ResponseWriter, r *http.Request) {
		deply(r.PostFormValue("groupid"))
		str, _ := json.Marshal(struct {
			Status string
		}{Status: "ok"})
		fmt.Fprintf(w, "%s\n", str)
	})

	http.HandleFunc("/rollback", func(w http.ResponseWriter, r *http.Request) {
		rollback(r.PostFormValue("groupid"), r.PostFormValue("reversion"))
		str, _ := json.Marshal(struct {
			Status string
		}{Status: "ok"})
		fmt.Fprintf(w, "%s\n", str)
	})

	s := &http.Server{
		Addr:           cfg.ListenHttp,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println(s.ListenAndServe())
	}()

	//start tcp server
	log.Printf("Tcp server start,listening at %s", cfg.ListenTcp)

	ln, err := net.Listen("tcp", cfg.ListenTcp)
	if err != nil {
		log.Println("Error listening:", err)
		os.Exit(1)
	}
	defer ln.Close()

	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			continue
		}
		log.Printf("Received new connection %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	if !checkAllowedConn(conn) {
		return
	} else {
		setClientStatus(conn, true)
	}
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Client closed", err.Error())
			break
		}
		// output message received
		log.Print(conn.RemoteAddr(), " -> Message Received:", message)
		//keepalive
		if strings.TrimSpace(message) == "PING" {
			message = "PONG"
		}
		// send new string back to client
		conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Println("Error writing to stream.", err)
			break
		}
	}
	setClientStatus(conn, false)
}

//check connection whether to allow
func checkAllowedConn(conn net.Conn) bool {
	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	for _, app := range cfg.Apps {
		for _, node := range app.Node {
			if node.Ip == ip {
				if node.Online {
					return false
				}
				return true
			}
		}
	}
	return false
}

//set client online or offline
func setClientStatus(conn net.Conn, online bool) {
	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	for key, app := range cfg.Apps {
		for k, node := range app.Node {
			if node.Ip == ip {
				cfg.Apps[key].Node[k].Online = online
				if online {
					cfg.Apps[key].Node[k].Conn = conn
				} else {
					cfg.Apps[key].Node[k].Conn = nil
				}
			}
		}
	}
}

//send a update message to the group nodes
func deply(groupid string) {
	for _, app := range cfg.Apps {
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
	for _, app := range cfg.Apps {
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
