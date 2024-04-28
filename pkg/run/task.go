package run

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
)

func dispatchTask(task task.Task, input libtask.TaskInput, targets host.Targets) libtask.TaskOutput {
	lenTargets := config.LenTargets(targets)
	var out libtask.TaskOutput
	for _, hosts := range targets {
		for i, h := range hosts {
			if !input.Quiet {
				fmt.Printf("  host: %s (protocol %s) (%d/%d)\n", h.Client.Address(), h.Client.Protocol(), i+1, lenTargets)
			}
			out = task.Run(input, h)
		}
	}
	return out
}

func matchHostsToTask(playHosts []string, targets host.Targets) host.Targets {
	taskTargets := host.Targets{}
	for alias, host := range targets {
		if slices.Contains(playHosts, alias) {
			taskTargets[alias] = host
		}
	}
	return taskTargets
}
