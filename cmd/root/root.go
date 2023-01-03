
package root

import (
	"github.com/sikalabs/gobble/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gobble",
	Short: "gobble, " + version.Version,
}
