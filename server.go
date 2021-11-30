package main

import (
	"bufio"
	"go-deploy/config"
	"go-deploy/ctrl"
	"log"
	"net"
	"net/http"
	"time"
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
		}
	}()

	for addr := range config.C.UniqAddr {
		go Ping(addr)
	}

	//block
	ticker := time.Tick(time.Second * 10)
	for {
		select {
		case <-ticker:
		}
	}
}

func Ping(addr string) {
	for {
		func() {
			//connect to this socket
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Println("Error connect to client:", err)
				return
			}

			//connect success
			setClientOnlineStatus(addr, true)

			//remote client closed
			defer func() {
				setClientOnlineStatus(addr, false)
				conn.Close()
			}()

			//read message from client
			go func(conn net.Conn) {
				defer conn.Close()
				for {
					message, err := bufio.NewReader(conn).ReadString('\n')
					if err != nil {
						log.Println("Client closed", err.Error())
						return
					}
					log.Print(conn.RemoteAddr(), " -> Message Received from client:", message)
				}
			}(conn)

			ticker := time.Tick(time.Second * 15)
			for {
				select {
				case <-ticker:
					conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
					_, err := conn.Write([]byte("PING\n"))
					if err != nil {
						log.Println("Error writing to stream:", err)
						return
					}
				}
			}
		}()
		time.Sleep(time.Second * 5)
	}
}

//set client online or offline
func setClientOnlineStatus(addr string, online bool) {
	for key, app := range config.C.Apps {
		for k, node := range app.Node {
			if node.Addr == addr {
				config.C.Apps[key].Node[k].Online = online
			}
		}
	}
}
