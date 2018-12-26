package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	PROTOCOL = "tcp"
	PORT2 = ":7777"
)

func main() {
	tcpConnection, err := net.Dial(PROTOCOL, PORT2)
	if err != nil {
		panic(err.Error())
	}

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("unable o read input: ", err.Error())
			break
		}

		// send to socket
		fmt.Fprintf(tcpConnection, text)

		// listen for reply
		for {
			buffer := make([]byte, 999999999)
			dataSize, err := tcpConnection.Read(buffer)
			if err != nil {
				fmt.Println("The connection has closed!")
				return
			}

			data := buffer[:dataSize]
			fmt.Println("Received message:\n", string(data))
			break
		}
	}
}