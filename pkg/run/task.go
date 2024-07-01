package run

import (
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/printer"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
	"sync"
)

func DispatchTask(task task.Task, input libtask.TaskInput, targets host.Targets) libtask.TaskOutput {
	lenTargets := config.LenTargets(targets)
	printer.GlobalPrinter.SetHostLength(lenTargets)
	var out libtask.TaskOutput
	for _, hosts := range targets {
		for _, h := range hosts {
			printer.GlobalPrinter.PrintHost(h.Client.Address(), h.Client.Protocol())
			out = task.Run(input, h)
		}
	}
	return out
}

func DispatchTaskP(task task.Task, input libtask.TaskInput, targets host.Targets) libtask.TaskOutput {
	lenTargets := config.LenTargets(targets)
	printer.GlobalPrinter.SetHostLength(lenTargets)
	results := make(chan libtask.TaskOutput, lenTargets)
	var wg sync.WaitGroup

	for _, hosts := range targets {
		for i, h := range hosts {
			wg.Add(1) // Increment the WaitGroup counter.
			go func(h *host.Host, i int) {
				defer wg.Done()
				results <- task.Run(input, h) // Run task and send the result to the channel.
				printer.GlobalPrinter.PrintHost(h.Client.Address(), h.Client.Protocol())
			}(h, i)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Aggregate results from the channel.
	var finalOutput libtask.TaskOutput
	for result := range results {
		if result.Error != nil {
			finalOutput.Error = result.Error
			break
		}
	}

	return finalOutput
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
