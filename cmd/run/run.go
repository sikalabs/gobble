package version

import (
	"log"

	"github.com/sikalabs/gobble/cmd/root"
	"github.com/sikalabs/gobble/pkg/run"
	"github.com/spf13/cobra"
)

var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:     "run",
	Short:   "Run gobblefile.yml",
	Aliases: []string{"r"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := run.Run(FlagDryRun)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Dry run",
	)
}
