package cmd

import (
	_ "github.com/sikalabs/gobble/cmd/list_schema_versions"
	"github.com/sikalabs/gobble/cmd/root"
	_ "github.com/sikalabs/gobble/cmd/run"
	_ "github.com/sikalabs/gobble/cmd/utils"
	_ "github.com/sikalabs/gobble/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
