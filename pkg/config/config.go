package config

import "github.com/sikalabs/gobble/pkg/play"

type ConfigMeta struct {
	SchemaVersion int `yaml:"schema_version"`
}

type ConfigHost struct {
	SSHTarget string                 `yaml:"ssh_target"`
	Vars      map[string]interface{} `yaml:"vars"`
}

type GlobalConfig struct {
	NoStrictHostKeyChecking bool                   `yaml:"no_strict_host_key_checking"`
	Vars                    map[string]interface{} `yaml:"vars"`
}

type Config struct {
	Meta   ConfigMeta              `yaml:"meta"`
	Global GlobalConfig            `yaml:"global"`
	Hosts  map[string][]ConfigHost `yaml:"hosts"`
	Plays  []play.Play             `yaml:"plays"`
}
