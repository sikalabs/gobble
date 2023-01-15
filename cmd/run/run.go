package version

import (
	"log"

	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/run"
	"github.com/spf13/cobra"
)

var FlagDryRun bool
var FlagOnlyTags []string
var FlagConfigFilePath string

var Cmd = &cobra.Command{
	Use:     "run",
	Short:   "Run gobblefile.yml",
	Aliases: []string{"gobble", "r"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := run.Run(FlagConfigFilePath, FlagDryRun, FlagOnlyTags)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagConfigFilePath,
		"config",
		"c",
		"gobblefile.yml",
		"Path to config file",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Dry run",
	)
	Cmd.Flags().StringSliceVar(
		&FlagOnlyTags,
		"only-tag",
		[]string{},
		"Run only selected task by tags",
	)
}
