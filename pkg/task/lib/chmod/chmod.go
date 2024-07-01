package chmod

import (
	"fmt"
	"github.com/k0sproject/rig/v2/remotefs"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"io/fs"
	"strconv"
)

type Task struct {
	libtask.BaseTask
	Path string `yaml:"path"`
	Perm string `yaml:"perm"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	// Create a new remote FS instance
	rfs := remotefs.NewFS(host.Client)
	if taskInput.Sudo {
		rfs = remotefs.NewFS(host.Client.Sudo())
	}

	// Parse the file mode
	mode, err := strconv.ParseUint(t.Perm, 8, 32)
	if err != nil {
		err = fmt.Errorf("invalid file mode: %s", err)
		return libtask.TaskOutput{Error: err}
	}

	// Change the file mode
	err = rfs.Chmod(t.Path, fs.FileMode(mode))

	return libtask.TaskOutput{
		Error: err,
	}
}
