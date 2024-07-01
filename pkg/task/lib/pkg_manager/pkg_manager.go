package pkg_manager

import (
	"context"
	"fmt"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"time"
)

type Task struct {
	libtask.BaseTask
	Package string `yaml:"name"`
	State   string `yaml:"state"`
	Update  bool   `yaml:"update"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get the package manager
	pkgManager, err := host.Client.PackageManagerService.GetPackageManager()
	if taskInput.Sudo {
		pkgManager, err = host.Client.Sudo().GetPackageManager()
	}

	if err != nil {
		return libtask.TaskOutput{Error: err}
	}
	// update packages
	if t.Update {
		err := pkgManager.Update(ctx)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}
	}

	// install or remove the package
	if t.State == "" {
		t.State = "present"
	}
	if t.State == "present" {
		err := pkgManager.Install(ctx, t.Package)
		return libtask.TaskOutput{
			Error: err,
		}
	} else if t.State == "absent" {
		err := pkgManager.Remove(ctx, t.Package)
		return libtask.TaskOutput{
			Error: err,
		}
	}
	return libtask.TaskOutput{
		Error: fmt.Errorf("unknown state: %s", t.State),
	}
}
