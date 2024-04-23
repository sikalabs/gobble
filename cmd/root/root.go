package root

import (
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/version"
	"github.com/spf13/cobra"
)

var FlagVerbosity int
var Cmd = &cobra.Command{
	Use:   "gobble",
	Short: "gobble, " + version.Version,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the logger with the verbosity flag
		logger.InitCharmLogger(FlagVerbosity)
	},
}

func init() {
	// Add a persistent flag for verbosity to the root command
	Cmd.PersistentFlags().IntVarP(&FlagVerbosity,
		"verbosity",
		"v",
		1,
		"Set the logging verbosity (1=error, 2=warn, 3=info, 4=debug)")
}
