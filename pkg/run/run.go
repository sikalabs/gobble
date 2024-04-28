package run

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/play"
	"golang.org/x/exp/slices"
)

func RunFromFile(
	configFilePath string,
	dryRun bool,
	quietOutput bool,
	onlyTags []string,
	skipTags []string,
) error {
	c, err := config.ReadConfigFile(configFilePath)
	if err != nil {
		return err
	}

	return Run(c, dryRun, quietOutput, onlyTags, skipTags)
}

func Run(
	c *config.Config,
	dryRun bool,
	quietOutput bool,
	onlyTags []string,
	skipTags []string,
) error {
	if c.Meta.SchemaVersion != 4 {
		return fmt.Errorf("unsupported schema version: %d", c.Meta.SchemaVersion)
	}

	//Initialize host connections
	targets, err := host.InitializeHosts(c.RigHosts, c.HostsAliases)
	if err != nil {
		logger.Log.Fatal(err)
	}

	c.AllPlays = []play.Play{}

	// Add Includes Before
	for _, includePlays := range c.IncludePlaysBefore {
		plays, err := play.GetPlaysFromIncludePlays(includePlays)
		if err != nil {
			logger.Log.Fatal(err)
		}
		c.AllPlays = append(c.AllPlays, plays...)
	}

	// Add Plays from Gobblefile
	c.AllPlays = append(c.AllPlays, c.Plays...)

	// Add Includes After
	for _, includePlays := range c.IncludePlaysAfter {
		plays, err := play.GetPlaysFromIncludePlays(includePlays)
		if err != nil {
			logger.Log.Fatal(err)
		}
		c.AllPlays = append(c.AllPlays, plays...)
	}

	lenPlays := config.LenPlays(c, onlyTags, skipTags)
	playI := 0
	for _, play := range c.AllPlays {
		skip := false
		for _, tag := range skipTags {
			if slices.Contains(play.Tags, tag) {
				skip = true
			}
		}
		if skip {
			continue
		}
		if len(onlyTags) > 0 {
			skip = true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
		}
		if skip {
			continue
		}
		playI++

		lenTasks := len(play.Tasks)
		taskI := 0
		for _, t := range play.Tasks {
			taskI++

			if !quietOutput {
				fmt.Printf("+ play: %s (%d/%d)\n", play.Name, playI, lenPlays)
				fmt.Printf("  task: %s (%d/%d)\n", t.GetName(), taskI, lenTasks)
				if play.Sudo {
					fmt.Printf("  sudo: %t\n", play.Sudo)
				}
			}

			taskTargets := matchHostsToTask(play.Hosts, targets)
			taskInput := libtask.TaskInput{
				Config:                  c,
				NoStrictHostKeyChecking: c.Global.NoStrictHostKeyChecking,
				Sudo:                    play.Sudo,
				Vars:                    c.Global.Vars,
				Dry:                     dryRun,
				Quiet:                   quietOutput,
			}
			out := DispatchTask(t, taskInput, taskTargets)
			if out.Error != nil {
				return out.Error
			}
			fmt.Println(``)

			return nil
		}
	}
	return nil
}
