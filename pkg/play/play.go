package play

import "github.com/sikalabs/gobble/pkg/task"

type Play struct {
	Name  string      `yaml:"name"`
	Sudo  bool        `yaml:"sudo"`
	Hosts []string    `yaml:"hosts"`
	Tasks []task.Task `yaml:"tasks"`
}
