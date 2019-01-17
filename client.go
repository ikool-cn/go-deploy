package main

import (
	"net"
	"bufio"
	"time"
	"log"
	"flag"
	"fmt"
	"os"
	"encoding/json"
	"strings"
	"go-deploy/helper"
)

var usage = `Usage: /data/pathto/client -s 127.0.0.1:8081`
var serverAddr *string

type Message struct {
	Type        string `json:"type"`
	Path        string `json:"path"`
	Action      string `json:"action"`
	Reversion   string `json:"reversion,omitempty"`
	BeforDeploy string `json:"befor_deploy"`
	AfterDeploy string `json:"after_deploy"`
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	serverAddr = flag.String("s", "127.0.0.1:8081", "server address")
	flag.Parse()
	if *serverAddr == "" {
		flag.Usage()
		os.Exit(1)
	}

	for {
		connectServer()
		time.Sleep(time.Second * 5)
	}
}

func connectServer() {
	// connect to this socket
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		log.Println("Error connect to server:", err)
		return
	}
	defer conn.Close()
	go handleServer(conn)

	ticker := time.Tick(time.Second * 15)
	for {
		select {
		case <-ticker:
			conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
			_, err := conn.Write([]byte("PING\n"))
			if err != nil {
				log.Println("Error writing to stream.", err)
				return
			}
		}
	}
}

func handleServer(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Server closed", err.Error())
			return
		}
		log.Print("Message received from server: " + message)
		if message != "PONG\n" {
			go processAction(message)
		}
	}
}

func processAction(message string) {
	msg := &Message{}
	err := json.Unmarshal([]byte(message), msg)
	if err != nil {
		log.Print("Json decode error: " + err.Error())
		return
	}

	var command string
	if msg.Action == "update" {
		command = fmt.Sprintf("cd %s && svn up", msg.Path)
	} else if msg.Action == "rollback" {
		command = fmt.Sprintf("cd %s && svn up -r %s", msg.Path, msg.Reversion)
	}

	if command != "" {
		if strings.TrimSpace(msg.BeforDeploy) != "" {
			helper.RunShell(msg.BeforDeploy)
		}
		helper.RunShell(command)
		if strings.TrimSpace(msg.AfterDeploy) != "" {
			helper.RunShell(msg.AfterDeploy)
		}
	}
}
