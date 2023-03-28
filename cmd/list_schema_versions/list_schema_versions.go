package list_schema_versions

import (
	"fmt"

	"github.com/sikalabs/gobble/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list-schema-versions",
	Short: "List of schema_version compatibilities",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Print(`List of schema_version compatibilities

x----------------x-----------------x
| schema_version | gobble versions |
x----------------x-----------------x
| 3              | v0.3.0 +        |
| 2              | v0.2.0          |
| 1              | v0.1.0          |
x----------------x-----------------x
`)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
