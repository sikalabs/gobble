package run

import (
	"context"
	"fmt"
	"github.com/k0sproject/rig/v2/cmd"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
	"time"
)

func dispatchTask(task task.Task, input libtask.TaskInput, targets Targets) libtask.TaskOutput {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var err error
	defer cancel()
	for _, hosts := range targets {
		for _, host := range hosts {
			fmt.Printf("Running on host: %s\n", host.Client.String())
			command := fmt.Sprintf("touch $HOME/%s", host.Client.Address())
			err = host.Client.ExecContext(ctx, command,
				cmd.LogInput(true),
				cmd.StreamOutput(),
				cmd.LogError(true),
				cmd.Logger(logger.Slog))
			if err != nil {
				logger.Log.Warn("task execution failed")
			}
		}
	}
	return libtask.TaskOutput{Error: err}
}

// func DispatchTask() {
//
//		for globalHostName, globalHost := range c.AllHosts {
//			for _, host := range globalHost {
//				if !slices.Contains(play.Hosts, globalHostName) {
//					continue
//				}
//
//				hostI++
//				//dispatchTask()
//
//				if !quietOutput {
//					fmt.Printf("+ play: %s (%d/%d)\n", play.Name, playI, lenPlays)
//					fmt.Printf("  task: %s (%d/%d)\n", t.Name, taskI, lenTasks)
//					if host.SSHPort == 0 {
//						fmt.Printf("  host: %s (%d/%d)\n", host.SSHTarget, hostI, lenHosts)
//					} else {
//						fmt.Printf("  host: %s (port %d) (%d/%d)\n", host.SSHTarget, host.SSHPort, hostI, lenHosts)
//					}
//					if play.Sudo {
//						fmt.Printf("  sudo: %t\n", play.Sudo)
//					}
//				}
//				taskInput := libtask.TaskInput{
//					SSHTarget:               host.SSHTarget,
//					SSHPort:                 host.SSHPort,
//					SSHPassword:             host.SSHPassword,
//					SSHOptions:              host.SSHOptions,
//					SudoPassword:            host.SudoPassword,
//					Config:                  c,
//					NoStrictHostKeyChecking: c.Global.NoStrictHostKeyChecking,
//					Sudo:                    play.Sudo,
//					Vars:                    mergeMaps(c.Global.Vars, host.Vars),
//					Dry:                     dryRun,
//					Quiet:                   quietOutput,
//				}
//				out := task.Run(taskInput, t)
//				if out.Error != nil {
//					return out.Error
//				}
//				fmt.Println(``)
//			}
//		}
//	}
func matchHostsToTask(playHosts []string, targets Targets) Targets {
	taskTargets := Targets{}
	for alias, host := range targets {
		if slices.Contains(playHosts, alias) {
			taskTargets[alias] = host
		}
	}
	return taskTargets
}
