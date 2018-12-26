package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGetUID(t *testing.T) {
	uidCmd := exec.Command("bash", "-c", "id -u")
	uidBytes, err := uidCmd.Output()
	uidFromCommand := strings.Trim(string(uidBytes), "\n")

	//usernameCmd := exec.Command("bash", "-c", "id -u -nr")
	//usernameBytes, err := usernameCmd.Output()
	//fmt.Println(string(usernameBytes), err)

	//using os package
	osUID := fmt.Sprintf("%v", os.Getuid())


	uid, _ := getUID()
	if err != nil || uid != uidFromCommand || uid != osUID {
		t.Fatalf("GetUID failed: %s == %s", uid, string(uidBytes))
	}
}

func TestGetPID(t *testing.T) {
	//p, _ := os.Executable()
	pidCmd := exec.Command("bash", "-c", fmt.Sprintf("ps aux | grep %s | awk 'NR==3 {print $2}'", os.Args[0]))
	pidBytes, err := pidCmd.Output()

	if err != nil || getPID() != strings.TrimSpace(string(pidBytes)) {
		t.Fatalf("GetPID failed: %s == %s", getPID(), strings.TrimSpace(string(pidBytes)))
	}
}
