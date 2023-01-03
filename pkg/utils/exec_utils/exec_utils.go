package exec_utils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
)

func Exec(taskInput libtask.TaskInput, cmd string, args ...string) error {
	if taskInput.Dry {
		fmt.Println(strings.Join(append([]string{cmd}, args...), " "))
		return nil
	}
	c := exec.Command(cmd, args...)
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr
	err := c.Run()
	return err
}
