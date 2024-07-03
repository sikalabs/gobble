package config

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/play"
)

type ConfigMeta struct {
	SchemaVersion int `yaml:"schema_version"`
}

type GlobalConfig struct {
	Vars map[string]interface{} `yaml:"vars"`
}

type Config struct {
	Meta               ConfigMeta                    `yaml:"meta"`
	Global             GlobalConfig                  `yaml:"global"`
	RigHosts           map[string][]*host.HostConfig `yaml:"hosts"`
	HostsAliases       map[string][]string           `yaml:"hosts_aliases"`
	Plays              []play.Play                   `yaml:"plays"`
	IncludePlaysBefore []play.InludePlays            `yaml:"include_plays_before"`
	IncludePlaysAfter  []play.InludePlays            `yaml:"include_plays_after"`

	AllPlays []play.Play
}
