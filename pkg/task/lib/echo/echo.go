package echo

import (
	"bytes"
	"context"
	"github.com/k0sproject/rig/v2/cmd"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/utils"
	"text/template"
	"time"

	"github.com/sikalabs/gobble/pkg/libtask"
)

type Task struct {
	libtask.BaseTask        // Embed BaseTask
	Message          string `yaml:"message"`
}

func (t Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Parse the message template
	tmpl, err := template.New("cmd").Parse(t.Message)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	// Execute the message template
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

	// Determine the client to use
	client := host.Client
	if taskInput.Sudo {
		client = host.Client.Sudo()
	}

	// execute the echo command with the message
	err = client.ExecContext(ctx, "echo "+buf.String(),
		cmd.LogInput(true),
		cmd.StreamOutput(),
		cmd.LogError(true),
		cmd.Logger(logger.Slog))
	if err != nil {
		logger.Log.Warnf("host: '%s' task: '%s' failed", client.Address(), t.GetName())
	}
	return libtask.TaskOutput{
		Error: err,
	}
}
