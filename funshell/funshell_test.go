package funshell_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/jthomperoo/test-exec-command-golang/funshell"
)

const (
	testStdoutValue = "test fun value!"
)

func TestShellCommand_Success(t *testing.T) {
	stdout, err := funshell.ShellCommand(fakeExecCommandSuccess)
	if err != nil {
		t.Error(err)
		return
	}

	stdoutStr := stdout.String()
	if stdoutStr != testStdoutValue {
		t.Errorf("stdout mismatch:\n%s\n vs \n%s", stdoutStr, testStdoutValue)
	}
}

func TestShellCommand_Fail(t *testing.T) {
	_, err := funshell.ShellCommand(fakeExecCommandFailure)
	if err == nil {
		t.Errorf("Expected error due to shell command exiting with non-zero exit code")
	}
}

// TestShellProcessSuccess is a method that is called as a substitute for a shell command,
// the GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is
// skipped.
func TestShellProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, testStdoutValue)
	os.Exit(0)
}

func TestShellProcessFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}

// fakeExecCommandSuccess is a function that initialises a new exec.Cmd, one which will
// simply call TestShellProcessSuccess rather than the command it is provided. It will
// also pass through the command and its arguments as an argument to TestShellProcessSuccess
func fakeExecCommandSuccess(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestShellProcessSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func fakeExecCommandFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestShellProcessFail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}
