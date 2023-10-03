package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	// create on heap.
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("conn created.")
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
