package cmd

import (
	"github.com/sikalabs/gobble/cmd/root"
	_ "github.com/sikalabs/gobble/cmd/run"
	_ "github.com/sikalabs/gobble/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
