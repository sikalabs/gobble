package config

import (
	"github.com/k0sproject/rig/v2/protocol/openssh"
	"github.com/k0sproject/rig/v2/protocol/ssh"
)

type HostConfig struct {
	SSH    *ssh.Config            `yaml:"ssh,omitempty"`
	Opensh *openssh.Config        `yaml:"openssh,omitempty"`
	Local  bool                   `yaml:"local,omitempty"`
	Vars   map[string]interface{} `yaml:"vars"`
}
