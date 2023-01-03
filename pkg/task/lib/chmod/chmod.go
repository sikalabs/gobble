package chmod

import (
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskChmod struct {
	Path string `yaml:"path"`
	Perm string `yaml:"perm"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskChmod,
) libtask.TaskOutput {
	err := exec_utils.Exec(
		taskInput,
		"ssh", taskInput.SSHTarget,
		"chmod", taskParams.Perm, taskParams.Path,
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
