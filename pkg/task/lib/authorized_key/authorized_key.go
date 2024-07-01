package authorized_key

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/k0sproject/rig/v2/remotefs"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"os"
	"path"
)

type Task struct {
	libtask.BaseTask
	Key string `yaml:"key"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	// Create a new FS instance
	rfs := remotefs.NewFS(host.Client)
	if taskInput.Sudo {
		rfs = remotefs.NewFS(host.Client.Sudo())
	}

	if rfs.FileExist("~/.ssh/authorized_keys") {
	}
	sshDir := path.Join(rfs.UserHomeDir(), ".ssh")
	authKeysFile := path.Join(sshDir, "authorized_keys")

	if !rfs.FileExist(sshDir) {
		err := rfs.MkdirAll(sshDir, 0755)
		if err != nil {
			return libtask.TaskOutput{Error: fmt.Errorf("failed to create directory: %w", err)}
		}
	}

	// Read the authorized_keys file
	data, err := rfs.ReadFile(authKeysFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			data = []byte{} // If the file doesn't exist, start with empty data
		} else {
			return libtask.TaskOutput{Error: err}
		}
	}

	// Check if the key is already in the file
	if bytes.Contains(data, []byte(t.Key)) {
		return libtask.TaskOutput{Error: nil} // Key already exists, no need to append
	}

	// Append the key to the authorized_keys file
	data = append(data, []byte("\n"+t.Key+"\n")...)
	if err := rfs.WriteFile(authKeysFile, data, 0644); err != nil {
		return libtask.TaskOutput{Error: err}
	}

	return libtask.TaskOutput{Error: nil}
}
