package host

import (
	"github.com/k0sproject/rig/v2"
	"github.com/k0sproject/rig/v2/protocol/openssh"
	"github.com/k0sproject/rig/v2/protocol/ssh"
	"github.com/k0sproject/rig/v2/remotefs"
)

type HostConfig struct {
	SSH    *SSHConfig             `yaml:"ssh,omitempty"`
	Opensh *openssh.Config        `yaml:"openssh,omitempty"`
	Local  bool                   `yaml:"local,omitempty"`
	Vars   map[string]interface{} `yaml:"vars"`
}

type Host struct {
	Client *rig.Client
	Fs     remotefs.FS
	Vars   map[string]interface{}
}

type SSHConfig struct {
	ssh.Config
	Password string `yaml:"password"`
}
