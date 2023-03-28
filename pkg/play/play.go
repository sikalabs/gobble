package play

import "github.com/sikalabs/gobble/pkg/task"

type Play struct {
	Name  string      `yaml:"name"`
	Sudo  bool        `yaml:"sudo"`
	Tags  []string    `yaml:"tags"`
	Hosts []string    `yaml:"hosts"`
	Tasks []task.Task `yaml:"tasks"`
}
