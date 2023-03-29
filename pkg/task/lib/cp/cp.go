package cp

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskCp struct {
	// Copy from local to remote
	LocalSrc  string `yaml:"local_src"`
	RemoteDst string `yaml:"remote_dst"`
	// Copy from remote to local
	RemoteSrc string `yaml:"remote_src"`
	LocalDst  string `yaml:"local_dst"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskCp,
) libtask.TaskOutput {
	if taskParams.LocalSrc != "" && taskParams.RemoteDst != "" {
		err := exec_utils.SCP(taskInput, taskParams.LocalSrc, taskParams.RemoteDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if taskParams.RemoteSrc != "" && taskParams.LocalDst != "" {
		err := exec_utils.SCPRemoteToLocal(taskInput, taskParams.RemoteSrc, taskParams.LocalDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else {
		return libtask.TaskOutput{
			Error: fmt.Errorf("invalid cp task params"),
		}
	}
}
