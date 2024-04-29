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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pkgManager, err := host.Client.PackageManagerService.GetPackageManager()
	if err != nil {
		return libtask.TaskOutput{Error: err}
	}

	if t.Update {
		err := pkgManager.Update(ctx)
		if err != nil {
			return libtask.TaskOutput{
				Error: err,
			}
		}
	}

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
