package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
	"trusting-social/tcp-server/model"
)

const (
	PROTOCOL = "tcp"
	PORT1 = ":6777"
)

func main()  {
	tcpConnection, err := net.Dial(PROTOCOL, PORT1)
	if err != nil {
		panic(err.Error())
	}
	for i := 0;i <= 100; i++ {
		user := &model.User{
			ID: fmt.Sprintf("%v", rand.Int()),
			Name: PROTOCOL,
			Age: fmt.Sprintf("%v", rand.Int()),
		}
		data, err := json.Marshal(user)
		if err != nil {
			panic(err.Error())
		}
		n, err := tcpConnection.Write(data)
		fmt.Println(n, err)
		time.Sleep(1*time.Second)
	}
}
