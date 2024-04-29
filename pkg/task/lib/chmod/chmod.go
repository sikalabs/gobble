package chmod

import (
	"fmt"
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

	mode, err := strconv.ParseUint(t.Perm, 8, 32)
	if err != nil {
		err = fmt.Errorf("invalid file mode: %s", err)
		return libtask.TaskOutput{Error: err}
	}

	err = host.Fs.Chmod(t.Path, fs.FileMode(mode))

	return libtask.TaskOutput{
		Error: err,
	}
}
