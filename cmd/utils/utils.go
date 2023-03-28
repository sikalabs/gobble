package utils

import (
	"github.com/sikalabs/gobble/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "utils",
	Short: "Built-in utils",
	Args:  cobra.NoArgs,
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
