package host

import (
	"github.com/k0sproject/rig/v2"
	"github.com/k0sproject/rig/v2/protocol/openssh"
	"github.com/k0sproject/rig/v2/protocol/ssh"
)

type HostConfig struct {
	SSH          *SSHConfig             `yaml:"ssh,omitempty"`
	Opensh       *openssh.Config        `yaml:"openssh,omitempty"`
	Local        bool                   `yaml:"local,omitempty"`
	Vars         map[string]interface{} `yaml:"vars"`
	SudoPassword string                 `yaml:"sudo_password,omitempty"`
}

type Host struct {
	Client *rig.Client
	Vars   map[string]interface{}
}

type SSHConfig struct {
	ssh.Config `yaml:",inline"`
	Password   string `yaml:"password,omitempty"`
}
