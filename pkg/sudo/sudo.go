package sudo

import (
	"errors"
	"fmt"
	"github.com/k0sproject/rig/v2/cmd"
	"github.com/k0sproject/rig/v2/plumbing"
	"github.com/k0sproject/rig/v2/sh/shellescape"
	"github.com/k0sproject/rig/v2/sudo"
)

var (
	// ErrNoSudo is returned when no supported sudo method is found.
	ErrNoSudo = errors.New("no supported sudo method found")
)

// NewSudoProviderWithPass creates a new sudo provider configured with a sudo password.
func NewSudoProviderWithPass(password string) *sudo.Provider {
	provider := plumbing.NewProvider[cmd.Runner, cmd.Runner](ErrNoSudo)
	provider.Register(func(c cmd.Runner) (cmd.Runner, bool) {
		if c.IsWindows() {
			return nil, false
		}
		decorator := func(command string) string {
			return SudoPass(command, password)
		}
		return cmd.NewExecutor(c, decorator), true
	})
	return provider
}

// SudoPass is a DecorateFunc that will wrap the given command in a sudo call.
func SudoPass(cmd string, pass string) string {
	return fmt.Sprintf(`echo %s | sudo -S -- "${SHELL-sh}" -c %s`, shellescape.Quote(pass), shellescape.Quote(cmd))

}
