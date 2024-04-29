package authorized_key

import (
	"bytes"
	"errors"
	"fmt"
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

	if host.Fs.FileExist("~/.ssh/authorized_keys") {
	}
	sshDir := path.Join(host.Fs.UserHomeDir(), ".ssh")
	authKeysFile := path.Join(sshDir, "authorized_keys")

	if !host.Fs.FileExist(sshDir) {
		err := host.Fs.MkdirAll(sshDir, 0755)
		if err != nil {
			return libtask.TaskOutput{Error: fmt.Errorf("failed to create directory: %w", err)}
		}
	}

	// Read the authorized_keys file
	data, err := host.Fs.ReadFile(authKeysFile)
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
	if err := host.Fs.WriteFile(authKeysFile, data, 0644); err != nil {
		return libtask.TaskOutput{Error: err}
	}

	return libtask.TaskOutput{Error: nil}
}
