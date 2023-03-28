package yaml_merge

import (
	parent_command "github.com/sikalabs/gobble/cmd/utils"
	"github.com/sikalabs/gobble/pkg/utils/yaml_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "yaml-merge <YAML> <YAML> ...",
	Short: "Merge yaml files",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		yaml_utils.MergeYAMLs(args)
	},
}

func init() {
	parent_command.Cmd.AddCommand(Cmd)
}
