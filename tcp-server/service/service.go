package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"trusting-social/tcp-server/model"
	"trusting-social/tcp-server/repository"
)

var (
	dataShelf = make([]*model.User, 0)
)

func AcceptJSON(server net.Conn) error {
	for {
		buffer := make([]byte, 9999)
		dataSize, err := server.Read(buffer)
		if err != nil {
			return err
		}

		data := buffer[:dataSize]
		fmt.Println("JSON Server: Received message: ", string(data))

		// bufio can be used instead of above lines
		//data, _ := bufio.NewReader(server).ReadString('\n')

		var user *model.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			return err
		}

		//append in memory/buffered storage
		dataShelf = append(dataShelf, user)
		go func(user *model.User) {
			err = repository.Write(user)
			if err != nil {
				fmt.Println("error while writing to file:: ", err.Error())
			}
		}(user)
	}
}

func AcceptBuffer(server net.Conn) error {
	for {
		buffer := make([]byte, 9999)
		dataSize, err := server.Read(buffer)
		if err != nil {
			return err
		}

		data := buffer[:dataSize]
		fmt.Println("Buffer Server: Received message: ", string(data))

		if len(dataShelf) > 0 {
			b, err := json.Marshal(dataShelf)
			if err != nil {
				return err
			}
			server.Write(b)
		} else {
			b, err := ioutil.ReadFile(repository.FILE_PATH)
			if err != nil {
				return err
			}

			server.Write(b)
		}
	}
}
