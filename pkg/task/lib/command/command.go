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
	err := exec_utils.SSH(
		taskInput,
		taskParams.Cmd,
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
