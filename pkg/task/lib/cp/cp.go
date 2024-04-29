package cp

import (
	"fmt"
	"github.com/k0sproject/rig/v2/remotefs"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/utils"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/template_utils"
)

type Task struct {
	libtask.BaseTask
	// Copy from local to remote
	LocalSrc  string `yaml:"local_src"`
	RemoteDst string `yaml:"remote_dst"`
	// Copy from remote to local
	RemoteSrc string `yaml:"remote_src"`
	LocalDst  string `yaml:"local_dst"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	vars := map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   utils.MergeMaps(taskInput.Vars, host.Vars),
	}

	if t.LocalSrc != "" && t.RemoteDst != "" {
		// Render TaskCp.LocalSrc string
		localSrc, err := template_utils.RenderTemplateToString(
			t.LocalSrc, "TaskCp.LocalSrc", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Render TaskCp.RemoteDst string
		remoteDst, err := template_utils.RenderTemplateToString(
			t.RemoteDst, "TaskCp.RemoteDst", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Upload
		err = remotefs.Upload(host.Fs, localSrc, remoteDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if t.RemoteSrc != "" && t.LocalDst != "" {
		// Render TaskCp.RemoteSrc string
		remoteSrc, err := template_utils.RenderTemplateToString(
			t.RemoteSrc, "TaskCp.RemoteSrc", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Render TaskCp.LocalDst string
		localDst, err := template_utils.RenderTemplateToString(
			t.LocalDst, "TaskCp.LocalDst", vars,
		)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}

		// Download
		err = remotefs.Download(host.Fs, remoteSrc, localDst)
		return libtask.TaskOutput{
			Error: err,
		}
	} else {
		return libtask.TaskOutput{
			Error: fmt.Errorf("invalid cp task params"),
		}
	}
}
