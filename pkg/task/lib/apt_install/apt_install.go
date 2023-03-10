package apt_install

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskAptInstall struct {
	Name  string `yaml:"name"`
	State string `yaml:"state"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskAptInstall,
) libtask.TaskOutput {
	var err error
	if taskParams.State == "" {
		taskParams.State = "present"
	}
	if taskParams.State == "present" {
		err = exec_utils.Exec(
			taskInput,
			"ssh", taskInput.SSHTarget,
			"apt-get", "install", "-y", "--no-install-recommends",
			taskParams.Name,
		)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if taskParams.State == "absent" {
		err = exec_utils.Exec(
			taskInput,
			"ssh", taskInput.SSHTarget,
			"apt-get", "purge", "-y",
			taskParams.Name,
		)
		return libtask.TaskOutput{
			Error: err,
		}
	}
	return libtask.TaskOutput{
		Error: fmt.Errorf("unknown state: %s", taskParams.State),
	}
}
