package config

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/play"
)

type ConfigMeta struct {
	SchemaVersion int `yaml:"schema_version"`
}

type ConfigHost struct {
	SSHTarget    string                 `yaml:"ssh_target"`
	SSHPort      int                    `yaml:"ssh_port"`
	SSHPassword  string                 `yaml:"ssh_password"`
	SSHOptions   []string               `yaml:"ssh_options"`
	SudoPassword string                 `yaml:"sudo_password"`
	Vars         map[string]interface{} `yaml:"vars"`
}

type GlobalConfig struct {
	NoStrictHostKeyChecking bool                   `yaml:"no_strict_host_key_checking"`
	Vars                    map[string]interface{} `yaml:"vars"`
}

type Config struct {
	Meta               ConfigMeta                    `yaml:"meta"`
	Global             GlobalConfig                  `yaml:"global"`
	Hosts              map[string][]ConfigHost       `yaml:"hosts"`
	RigHosts           map[string][]*host.HostConfig `yaml:"rig_hosts"`
	HostsAliases       map[string][]string           `yaml:"hosts_aliases"`
	Plays              []play.Play                   `yaml:"plays"`
	IncludePlaysBefore []play.InludePlays            `yaml:"include_plays_before"`
	IncludePlaysAfter  []play.InludePlays            `yaml:"include_plays_after"`

	AllHosts map[string][]ConfigHost
	AllPlays []play.Play
}
