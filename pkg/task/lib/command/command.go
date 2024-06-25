package command

import (
	"bytes"
	"context"
	"github.com/k0sproject/rig/v2/cmd"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/utils"
	"text/template"
	"time"
)

type Task struct {
	libtask.BaseTask        // Embed BaseTask
	Cmd              string `yaml:"cmd"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	tmpl, err := template.New("cmd").Parse(t.Cmd)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   utils.MergeMaps(taskInput.Vars, host.Vars),
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	client := host.Client
	if taskInput.Sudo {
		client = host.Client.Sudo()
	}

	err = client.ExecContext(ctx, buf.String(),
		cmd.LogInput(true),
		cmd.StreamOutput(),
		cmd.LogError(true),
		cmd.Logger(logger.Slog))
	if err != nil {
		logger.Log.Warnf("host: '%s' task: '%s' failed", host.Client.Address(), t.GetName())
	}
	return libtask.TaskOutput{
		Error: err,
	}
}
