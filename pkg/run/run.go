package run

import (
	"fmt"
	"io/ioutil"

	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task"
	"gopkg.in/yaml.v2"
)

func Run(dryRun bool) error {
	c := config.Config{}

	buf, err := ioutil.ReadFile("gobblefile.yml")
	if err != nil {
		return err
	}
	_ = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	if c.Meta.SchemaVersion != 1 {
		return fmt.Errorf("unsupported schema version: %d", c.Meta.SchemaVersion)
	}

	for _, play := range c.Plays {
		for _, t := range play.Tasks {
			for _, host := range c.Hosts[play.Hosts] {
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

	return nil
}
