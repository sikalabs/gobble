package version

import (
	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/run"
	"github.com/spf13/cobra"
)

var FlagDryRun bool
var FlagQuietOutput bool
var FlagOnlyTags []string
var FlagSkipTags []string
var FlagConfigFilePath string

var Cmd = &cobra.Command{
	Use:     "run",
	Short:   "Run gobblefile.yml",
	Aliases: []string{"gobble", "r"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := run.RunFromFile(FlagConfigFilePath, FlagDryRun, FlagQuietOutput, FlagOnlyTags, FlagSkipTags)
		if err != nil {
			logger.Log.Fatal(err)
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
		"Path to config file, \"-\" for STDIN",
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
	Cmd.Flags().StringSliceVar(
		&FlagSkipTags,
		"skip-tag",
		[]string{},
		"Skip only selected plays by tags",
	)
	Cmd.Flags().BoolVarP(
		&FlagQuietOutput,
		"quiet",
		"q",
		false,
		"Quiet output",
	)
}
