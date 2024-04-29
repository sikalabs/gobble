package ping

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/run"
	ping "github.com/sikalabs/gobble/pkg/task/lib/ping"
	"os"

	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/spf13/cobra"
)

var FlagConfigFilePath string

// TODO Reimplement ping as a task

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping all hosts from gobblefile",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		shellReturnCode := 0
		conf, err := config.ReadConfigFile(FlagConfigFilePath)
		if err != nil {
			logger.Log.Fatal(err)
		}

		if conf.Meta.SchemaVersion != 4 {
			logger.Log.Fatalf("unsupported schema version: %d", conf.Meta.SchemaVersion)
		}

		//Initialize host connections
		targets, err := host.InitializeHosts(conf.RigHosts, conf.HostsAliases)
		if err != nil {
			logger.Log.Fatal(err)
		}

		task := ping.Task{
			BaseTask: libtask.BaseTask{
				Name: "ping all hosts",
			},
		}
		ti := libtask.TaskInput{
			Config:                  conf,
			NoStrictHostKeyChecking: conf.Global.NoStrictHostKeyChecking,
			Sudo:                    false,
			Vars:                    conf.Global.Vars,
			Dry:                     false,
			Quiet:                   false,
		}

		out := run.DispatchTask(&task, ti, targets)
		isOK := out.Error == nil
		if isOK {
			logger.Log.Printf("Hosts are OK")
		} else {
			logger.Log.Errorf("Hosts are not reachable")
			shellReturnCode = 1
		}
		os.Exit(shellReturnCode)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagConfigFilePath,
		"config",
		"c",
		"gobblefile.yml",
		"Path to config file, \"-\" for STDIN",
	)
}
