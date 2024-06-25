package host

import (
	"context"
	"fmt"
	"github.com/k0sproject/rig/v2"
	"github.com/k0sproject/rig/v2/protocol"
	"github.com/k0sproject/rig/v2/protocol/localhost"
	"github.com/k0sproject/rig/v2/protocol/openssh"
	"github.com/k0sproject/rig/v2/protocol/ssh"
	"github.com/k0sproject/rig/v2/remotefs"
	"github.com/sikalabs/gobble/pkg/logger"
	"github.com/sikalabs/gobble/pkg/sudo"
	go_ssh "golang.org/x/crypto/ssh"
	"time"
)

type Targets map[string][]*Host

// InitializeHosts initializes hosts from config
func InitializeHosts(rigHosts map[string][]*HostConfig, hostAliases map[string][]string) (map[string][]*Host, error) {
	initializedHosts := make(map[string][]*Host)
	for globalHostName, globalHost := range rigHosts {
		for _, host := range globalHost {
			ihost, err := setupHost(host)
			if err != nil {
				logger.Log.Fatalf("failed to setup host %s: %s", globalHostName, err)
			}
			initializedHosts[globalHostName] = append(initializedHosts[globalHostName], ihost)
		}
	}
	// Map hosts to aliases
	targets := mapHostsToAlias(initializedHosts, hostAliases)
	return targets, nil
}

// setupHost configures and connects to a single host
func setupHost(hostConfig *HostConfig) (*Host, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := createConnection(hostConfig)
	if err != nil {
		return nil, err
	}
	client, err := rig.NewClient(rig.WithConnection(conn), rig.WithLogger(logger.Slog))
	if hostConfig.SudoPassword != "" {
		logger.Log.Warn("Using sudo with password is deprecated, set up passwordless along with ssh keys")
		client, err = rig.NewClient(rig.WithConnection(conn), rig.WithLogger(logger.Slog), rig.WithSudoProvider(sudo.NewSudoProviderWithPass(hostConfig.SudoPassword)))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create rig client: %w", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return &Host{Client: client, Vars: hostConfig.Vars, Fs: remotefs.NewFS(client)}, nil
}

// createConnection creates a connection to a host
func createConnection(hostConfig *HostConfig) (protocol.Connection, error) {

	if hostConfig.Local {

		return localhost.NewConnection()

	} else if hostConfig.SSH != nil {

		cfg := *hostConfig.SSH
		//handle password auth
		if cfg.Password != "" {
			logger.Log.Warn("Using ssh with password is deprecated, please use ssh keys")
			cfg.Config.AuthMethods = append(cfg.Config.AuthMethods, go_ssh.Password(cfg.Password))
			return ssh.NewConnection(cfg.Config)
		} else {
			return ssh.NewConnection(cfg.Config)
		}

	} else if hostConfig.Opensh != nil {

		cfg := *hostConfig.Opensh
		return openssh.NewConnection(cfg)

	} else {
		return nil, fmt.Errorf("no suitable connection found")
	}
}

func mapHostsToAlias(hosts map[string][]*Host, hostAliases map[string][]string) Targets {

	targets := hosts
	for name, aliases := range hostAliases {
		for _, alias := range aliases {
			targets[name] = append(targets[name], hosts[alias]...)
		}
	}
	return targets
}
