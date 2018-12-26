package main

import (
	"fmt"
	"net"
	"trusting-social/tcp-server/service"
)

const (
	PROTOCOL = "tcp"
	PORT1 = ":6777"
	PORT2 = ":7777"
)

func main() {
	tcpListener1, err1 := net.Listen(PROTOCOL, PORT1)
	if err1 != nil {
		panic(err1.Error())
	}

	tcpListener2, err2 := net.Listen(PROTOCOL, PORT2)
	if err1 != nil {
		panic(err2.Error())
	}

	go keepListeningJSON(tcpListener1)
	keepListening(tcpListener2)
}

func keepListeningJSON(listener net.Listener) {
	defer listener.Close()
	for {
		server, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error: ", err.Error())
		}
		service.AcceptJSON(server)
	}
}

func keepListening(listener net.Listener) {
	defer listener.Close()
	for {
		server, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error: ", err.Error())
		}
		service.AcceptBuffer(server)
	}
}