package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	UID, _ := getUID()
	PID := getPID()
	fmt.Printf("UID: %s, PID: %s", UID, PID)
}


func getUID() (string, string) {
	currentUser, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return currentUser.Uid, currentUser.Username
}

func getPID() (string) {
	return fmt.Sprintf("%v", os.Getpid())
}



