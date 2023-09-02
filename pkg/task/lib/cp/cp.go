package cp

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
	"github.com/sikalabs/gobble/pkg/utils/template_utils"
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
	vars := map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   taskInput.Vars,
	}

	if taskParams.LocalSrc != "" && taskParams.RemoteDst != "" {
		// Render TaskCp.LocalSrc string
		localSrc, err := template_utils.RenderTemplateToString(
			taskParams.LocalSrc, "TaskCp.LocalSrc", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Render TaskCp.RemoteDst string
		remoteDst, err := template_utils.RenderTemplateToString(
			taskParams.RemoteDst, "TaskCp.RemoteDst", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// do scp
		err = exec_utils.SCP(taskInput, localSrc, remoteDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if taskParams.RemoteSrc != "" && taskParams.LocalDst != "" {
		// Render TaskCp.RemoteSrc string
		remoteSrc, err := template_utils.RenderTemplateToString(
			taskParams.RemoteSrc, "TaskCp.RemoteSrc", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Render TaskCp.LocalDst string
		localDst, err := template_utils.RenderTemplateToString(
			taskParams.LocalDst, "TaskCp.LocalDst", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// do scp
		err = exec_utils.SCPRemoteToLocal(taskInput, remoteSrc, localDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else {
		return libtask.TaskOutput{
			Error: fmt.Errorf("invalid cp task params"),
		}
	}
}
