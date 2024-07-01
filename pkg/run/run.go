package run

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/play"
	"github.com/sikalabs/gobble/pkg/printer"
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

	// Filter plays by tags & set length
	filteredPlays := play.FilterPlays(c.AllPlays, onlyTags, skipTags)
	printer.GlobalPrinter.SetPlayLength(len(filteredPlays))

	// Run plays
	for _, p := range filteredPlays {
		printer.GlobalPrinter.PrintPlay(p.Name)
		printer.GlobalPrinter.SetTaskLength(len(p.Tasks))

		// filter tasks
		for _, t := range p.Tasks {
			printer.GlobalPrinter.PrintTask(t.GetName())
			taskTargets := matchHostsToTask(p.Hosts, targets)
			taskInput := libtask.TaskInput{
				Config: c,
				Sudo:   p.Sudo,
				Vars:   c.Global.Vars,
				Dry:    dryRun,
			}
			out := DispatchTaskP(t, taskInput, taskTargets)
			if out.Error != nil {
				return out.Error
			}
			fmt.Println(``)

		}
	}
	return nil
}
