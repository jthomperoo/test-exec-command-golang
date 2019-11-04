package funshell

import (
	"bytes"
	"os/exec"
)

type execContext = func(name string, arg ...string) *exec.Cmd

// ShellCommand prints out a fun shell command!
func ShellCommand(cmdContext execContext) (*bytes.Buffer, error) {
	cmd := cmdContext("echo", "'fun!'")

	// Set up byte buffers to read stdout
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &outb, nil
}
