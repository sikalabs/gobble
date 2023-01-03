package config

import "github.com/sikalabs/gobble/pkg/play"

type ConfigMeta struct {
	SchemaVersion int `yaml:"schema_version"`
}

type ConfigHost struct {
	SSHTarget string            `yaml:"ssh_target"`
	Vars      map[string]string `yaml:"vars"`
}

type Config struct {
	Meta  ConfigMeta              `yaml:"meta"`
	Hosts map[string][]ConfigHost `yaml:"hosts"`
	Plays []play.Play             `yaml:"plays"`
}
