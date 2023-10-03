package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	// user will be listening for incoming messages on this channel.
	ReadChan chan string
	conn     net.Conn
}

func newUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:     userAddr,
		Addr:     userAddr,
		ReadChan: make(chan string),
		conn:     conn,
	}

	go user.ListenComingMessage()

	return user
}

func (this *User) ListenComingMessage() {
	for {
		msg := <-this.ReadChan

		_, err := this.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("Conn.Write err: ", err)
		}
	}
}
