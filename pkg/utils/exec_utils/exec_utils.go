package exec_utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
)

func RawExec(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	err := c.Run()
	return err
}

func RawExecStdout(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	return err
}

func Exec(taskInput libtask.TaskInput, cmd string, args ...string) error {
	if taskInput.Dry {
		fmt.Println(strings.Join(append([]string{cmd}, args...), " "))
		return nil
	}
	err := RawExec(cmd, args...)
	return err
}
