package ping

import (
	"context"
	"fmt"
	"github.com/k0sproject/rig/v2/cmd"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/logger"
	"time"
)

type Task struct {
	libtask.BaseTask // Embed BaseTask
}

func (t Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	logger.Log.Printf("Pinging host %s using protocol: %s ...", host.Client.Address(), host.Client.Protocol())
	err := host.Client.ExecContext(ctx, fmt.Sprintf(" echo 'Hello from %s'", host.Client.Address()),
		cmd.LogInput(true),
		cmd.StreamOutput(),
		cmd.LogError(true),
		cmd.Logger(logger.Slog))
	if err == nil {
		logger.Log.Printf("Host %s is OK, protocol: %s", host.Client.Address(), host.Client.Protocol())
	} else {
		logger.Log.Errorf("Host %s is not reachable, protocol: %s", host.Client.Address(), host.Client.Protocol())
	}
	return libtask.TaskOutput{
		Error: err,
	}
}
