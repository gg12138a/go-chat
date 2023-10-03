package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// online user
	UserOnlineMap map[string]*User
	mapLock       sync.RWMutex

	// msg broadcast
	MsgChan chan string
}

func NewServer(ip string, port int) *Server {
	// create on heap.
	server := &Server{
		Ip:            ip,
		Port:          port,
		UserOnlineMap: make(map[string]*User),
		MsgChan:       make(chan string),
	}

	return server
}

func (this *Server) ListenMsg() {
	for {
		msg := <-this.MsgChan

		// broadcast to all user
		this.mapLock.RLock()
		for _, user := range this.UserOnlineMap {
			user.ReadChan <- msg
		}
		this.mapLock.RUnlock()
	}
}

func (this *Server) Broadcast(sendUser *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s: %s", sendUser.Addr, sendUser.Name, msg)
	this.MsgChan <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {

	user := newUser(conn)

	// user login
	this.mapLock.Lock()
	this.UserOnlineMap[user.Name] = user
	this.mapLock.Unlock()

	this.Broadcast(user, "Login")

}

func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	// close listening socket
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Listener.Close err: ", err)
		}
	}(listener)

	// init server
	go this.ListenMsg()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener.Accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}
