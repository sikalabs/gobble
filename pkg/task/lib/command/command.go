package command

import (
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskCommand struct {
	Cmd string `yaml:"cmd"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskCommand,
) libtask.TaskOutput {
	err := exec_utils.Exec(
		taskInput,
		"ssh", append([]string{taskInput.SSHTarget}, taskParams.Cmd)...)
	return libtask.TaskOutput{
		Error: err,
	}
}
