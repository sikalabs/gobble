package ping

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mohae/deepcopy"
	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/echo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var FlagConfigFilePath string

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping all hosts from gobblefile",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		shellReturnCode := 0
		conf, err := readConfigFile(FlagConfigFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		if conf.Meta.SchemaVersion != 3 {
			log.Fatalln(fmt.Errorf("unsupported schema version: %d", conf.Meta.SchemaVersion))
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
				Vars:                    mergeMaps(conf.Global.Vars, host.Vars),
				Dry:                     false,
				Quiet:                   false,
			}
			taskParams := echo.TaskEcho{
				Message: "ping",
			}

			fmt.Printf("Ping %s using SSH ...", host.SSHTarget)
			out := echo.Run(ti, taskParams)
			isOK := out.Error == nil
			if isOK {
				fmt.Printf(" OK\n")
			} else {
				fmt.Printf(" ERR\n")
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

func readConfigFile(configFilePath string) (config.Config, error) {
	var buf []byte
	var err error
	c := config.Config{}

	if configFilePath == "-" {
		// Read from stdin
		buf, err = ioutil.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return c, err
		}
	} else {
		// Read from file
		buf, err = ioutil.ReadFile(configFilePath)
		if err != nil {
			return c, err
		}
	}

	_ = yaml.Unmarshal(buf, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	if m1 == nil {
		m1 = make(map[string]interface{})
	}
	deepCopyM1 := deepcopy.Copy(m1).(map[string]interface{})
	for k, v := range m2 {
		deepCopyM1[k] = v
	}
	return deepCopyM1
}
