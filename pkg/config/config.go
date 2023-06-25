package config

import "github.com/sikalabs/gobble/pkg/play"

type ConfigMeta struct {
	SchemaVersion int `yaml:"schema_version"`
}

type ConfigHost struct {
	SSHTarget    string                 `yaml:"ssh_target"`
	SSHPassword  string                 `yaml:"ssh_password"`
	SSHOptions   []string               `yaml:"ssh_options"`
	SudoPassword string                 `yaml:"sudo_password"`
	Vars         map[string]interface{} `yaml:"vars"`
}

type GlobalConfig struct {
	NoStrictHostKeyChecking bool                   `yaml:"no_strict_host_key_checking"`
	Vars                    map[string]interface{} `yaml:"vars"`
}

type InludePlays struct {
	Source string `yaml:"source"`
}

type Config struct {
	Meta   ConfigMeta              `yaml:"meta"`
	Global GlobalConfig            `yaml:"global"`
	Hosts  map[string][]ConfigHost `yaml:"hosts"`
	Plays  []play.Play             `yaml:"plays"`

	IncludePlaysBefore []InludePlays `yaml:"include_plays_before"`
	IncludePlaysAfter  []InludePlays `yaml:"include_plays_after"`

	AllPlays []play.Play
}
