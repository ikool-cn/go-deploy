package main

import (
	"net"
	"os"
	"bufio"
	"strings"
	"time"
	"log"
	"net/http"
	"go-deploy/config"
	"go-deploy/ctrl"
)

func main() {
	//start http server
	log.Printf("Start http server listening %s", config.C.ListenHttp)
	http.HandleFunc("/", ctrl.Index)
	http.HandleFunc("/list", ctrl.List)
	http.HandleFunc("/deply", ctrl.Deply)
	http.HandleFunc("/rollback", ctrl.Rollback)
	http.HandleFunc("/showlog", ctrl.ShowLog)

	s := &http.Server{
		Addr:           config.C.ListenHttp,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		} else {

		}
	}()

	//start tcp server
	log.Printf("Start tcp server listening %s", config.C.ListenTcp)
	ln, err := net.Listen("tcp", config.C.ListenTcp)
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
	for _, app := range config.C.Apps {
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
	for key, app := range config.C.Apps {
		for k, node := range app.Node {
			if node.Ip == ip {
				config.C.Apps[key].Node[k].Online = online
				if online {
					config.C.Apps[key].Node[k].Conn = conn
				} else {
					config.C.Apps[key].Node[k].Conn = nil
				}
			}
		}
	}
}
