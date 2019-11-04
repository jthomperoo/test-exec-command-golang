package main

import (
	"log"
	"os/exec"

	"github.com/jthomperoo/test-exec-command-golang/funshell"
)

func main() {
	stdout, err := funshell.ShellCommand(exec.Command)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(stdout.String())
}
