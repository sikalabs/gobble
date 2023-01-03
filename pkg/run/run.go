package run

import (
	"fmt"
	"io/ioutil"

	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

func Run(
	configFilePath string,
	dryRun bool,
	onlyTags []string,
) error {
	c := config.Config{}

	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	_ = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	if c.Meta.SchemaVersion != 2 {
		return fmt.Errorf("unsupported schema version: %d", c.Meta.SchemaVersion)
	}

	for _, play := range c.Plays {
		for _, t := range play.Tasks {
			for globalHostName, globalHost := range c.Hosts {
				for _, host := range globalHost {
					if !slices.Contains(play.Hosts, globalHostName) {
						continue
					}

					if len(onlyTags) > 0 {
						skip := true
						for _, tag := range onlyTags {
							if slices.Contains(t.Tags, tag) {
								skip = false
							}
						}
						if skip {
							continue
						}
					}

					fmt.Println(`+ play:`, play.Name)
					fmt.Println(`  host:`, host.SSHTarget)
					fmt.Println(`  task:`, t.Name)
					taskInput := libtask.TaskInput{
						SSHTarget: host.SSHTarget,
						Config:    c,
						Vars:      host.Vars,
						Dry:       dryRun,
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
