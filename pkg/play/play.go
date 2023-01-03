package play

import "github.com/sikalabs/gobble/pkg/task"

type Play struct {
	Name  string      `yaml:"name"`
	Hosts []string    `yaml:"hosts"`
	Tasks []task.Task `yaml:"tasks"`
}
