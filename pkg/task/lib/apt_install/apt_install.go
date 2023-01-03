package apt_install

import (
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskAptInstall struct {
	Name string `yaml:"name"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskAptInstall,
) libtask.TaskOutput {
	err := exec_utils.Exec(
		taskInput,
		"ssh", taskInput.SSHTarget,
		"apt-get", "install", "-y", "--no-install-recommends",
		taskParams.Name,
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
