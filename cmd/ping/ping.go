package ping

import (
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/utils"
	"os"

	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/echo"
	"github.com/spf13/cobra"
)

var FlagConfigFilePath string

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

		if conf.Meta.SchemaVersion != 3 {
			logger.Log.Fatalf("unsupported schema version: %d", conf.Meta.SchemaVersion)
		}

		conf.AllHosts = conf.Hosts

		for hostAliasName, hostAliases := range conf.HostsAliases {
			for _, hostAlias := range hostAliases {
				conf.AllHosts[hostAliasName] = append(conf.AllHosts[hostAliasName], conf.Hosts[hostAlias]...)
			}
		}

		allHosts := map[string]config.ConfigHost{}

		for _, hosts := range conf.AllHosts {
			for _, host := range hosts {
				allHosts[host.SSHTarget] = host
			}
		}

		for _, host := range allHosts {
			ti := libtask.TaskInput{
				SSHTarget:               host.SSHTarget,
				SSHPassword:             host.SSHPassword,
				SSHOptions:              host.SSHOptions,
				SudoPassword:            host.SudoPassword,
				Config:                  conf,
				NoStrictHostKeyChecking: conf.Global.NoStrictHostKeyChecking,
				Sudo:                    false,
				Vars:                    utils.MergeMaps(conf.Global.Vars, host.Vars),
				Dry:                     false,
				Quiet:                   false,
			}
			taskParams := echo.TaskEcho{
				Message: "ping",
			}

			logger.Log.Printf("Ping host %s using SSH ...", host.SSHTarget)
			out := echo.Run(ti, taskParams)
			isOK := out.Error == nil
			if isOK {
				logger.Log.Printf("Host: %s OK", host.SSHTarget)
			} else {
				logger.Log.Errorf("Host: %s Not reachable", host.SSHTarget)
				shellReturnCode = 1
			}
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
