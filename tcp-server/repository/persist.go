package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"trusting-social/tcp-server/model"
)

const PATH  = "/Users/%s/trusting-social.json"

var FILE_PATH string
func init() {
	_, USERNAME :=  getUID()
	FILE_PATH = fmt.Sprintf(PATH, USERNAME)
}

func Write(user *model.User) error {
	if user == nil {
		return errors.New("User is NIL")
	}

	b, err := json.Marshal(user)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return err
	}

	// write to file, f.Write()
	_, err = file.Write([]byte(string(b)+"\n"))
	return err
}

func getUID() (string, string) {
	currentUser, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return currentUser.Uid, currentUser.Username
}
