package helper

import (
    "fmt"
	"os"
	"log"
	"os/exec"
	"strconv"
)

func CheckPrivileges() {
	isSudo := exec.Command("id", "-u")
	output, err := isSudo.Output()

	if err != nil {
		log.Fatal(err)
	}

	// 0 = root, 501 = non-root user
	i, err := strconv.Atoi(string(output[:len(output)-1]))
	if err != nil {
		log.Fatal(err)
	}

	if i != 0 {
		fmt.Println("This program must be run as root! (sudo)")
		os.Exit(0)
	}
}
