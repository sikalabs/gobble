package apt_install

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskAptInstall struct {
	Name   string `yaml:"name"`
	State  string `yaml:"state"`
	Update bool   `yaml:"update"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskAptInstall,
) libtask.TaskOutput {
	var err error
	if taskParams.Update {
		err = exec_utils.SSH(
			taskInput,
			"apt-get", "update",
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}
	}
	if taskParams.State == "" {
		taskParams.State = "present"
	}
	if taskParams.State == "present" {
		err = exec_utils.SSH(
			taskInput,
			"apt-get", "install", "-y", "--no-install-recommends",
			taskParams.Name,
		)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if taskParams.State == "absent" {
		err = exec_utils.SSH(
			taskInput,
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
