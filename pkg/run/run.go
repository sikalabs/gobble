package run

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mohae/deepcopy"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/play"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

func Run(
	configFilePath string,
	dryRun bool,
	onlyTags []string,
) error {
	c, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	if c.Meta.SchemaVersion != 3 {
		return fmt.Errorf("unsupported schema version: %d", c.Meta.SchemaVersion)
	}

	lenPlays := lenPlays(c, onlyTags)
	playI := 0
	for _, play := range c.Plays {
		if len(onlyTags) > 0 {
			skip := true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
			if skip {
				continue
			}
		}
		playI++

		lenTasks := len(play.Tasks)
		taskI := 0
		for _, t := range play.Tasks {
			taskI++
			lenHosts := lenHosts(c, play)
			hostI := 0
			for globalHostName, globalHost := range c.Hosts {
				for _, host := range globalHost {
					if !slices.Contains(play.Hosts, globalHostName) {
						continue
					}
					hostI++

					fmt.Printf("+ play: %s (%d/%d)\n", play.Name, playI, lenPlays)
					fmt.Printf("  host: %s (%d/%d)\n", host.SSHTarget, hostI, lenHosts)
					fmt.Printf("  sudo: %t\n", play.Sudo)
					fmt.Printf("  task: %s (%d/%d)\n", t.Name, taskI, lenTasks)
					taskInput := libtask.TaskInput{
						SSHTarget:               host.SSHTarget,
						Config:                  c,
						NoStrictHostKeyChecking: c.Global.NoStrictHostKeyChecking,
						Sudo:                    play.Sudo,
						Vars:                    mergeMaps(c.Global.Vars, host.Vars),
						Dry:                     dryRun,
					}
					out := task.Run(taskInput, t)
					if out.Error != nil {
						return out.Error
					}
					fmt.Println(``)
				}
			}
		}
	}

	return nil
}

func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	deepCopyM1 := deepcopy.Copy(m1).(map[string]interface{})
	for k, v := range m2 {
		deepCopyM1[k] = v
	}
	return deepCopyM1
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

func lenPlays(c config.Config, onlyTags []string) int {
	length := 0
	for _, play := range c.Plays {
		if len(onlyTags) > 0 {
			skip := true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
			if skip {
				continue
			}
		}
		length++
	}

	return length
}

func lenHosts(c config.Config, play play.Play) int {
	length := 0
	for globalHostName, globalHost := range c.Hosts {
		for _, _ = range globalHost {
			if !slices.Contains(play.Hosts, globalHostName) {
				continue
			}
			length++
		}
	}
	return length
}
