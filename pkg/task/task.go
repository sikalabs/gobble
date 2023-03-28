package task

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/apt_install"
	"github.com/sikalabs/gobble/pkg/task/lib/chmod"
	"github.com/sikalabs/gobble/pkg/task/lib/command"
	"github.com/sikalabs/gobble/pkg/task/lib/cp"
	"github.com/sikalabs/gobble/pkg/task/lib/print"
	"github.com/sikalabs/gobble/pkg/task/lib/template"
)

type Task struct {
	Name       string                     `yaml:"name"`
	AptInstall apt_install.TaskAptInstall `yaml:"apt_install"`
	Cp         cp.TaskCp                  `yaml:"cp"`
	Template   template.TaskTemplate      `yaml:"template"`
	Command    command.TaskCommand        `yaml:"command"`
	Chmod      chmod.TaskChmod            `yaml:"chmod"`
	Print      print.TaskPrint            `yaml:"print"`
}

func Run(
	taskInput libtask.TaskInput,
	task Task,
) libtask.TaskOutput {
	switch {
	case task.AptInstall.Name != "":
		return apt_install.Run(taskInput, task.AptInstall)
	case task.Cp.LocalSrc != "":
		return cp.Run(taskInput, task.Cp)
	case task.Template.Path != "":
		return template.Run(taskInput, task.Template)
	case task.Command.Cmd != "":
		return command.Run(taskInput, task.Command)
	case task.Chmod.Path != "":
		return chmod.Run(taskInput, task.Chmod)
	case task.Print.Template != "":
		return print.Run(taskInput, task.Print)
	}
	return libtask.TaskOutput{
		Error: fmt.Errorf("task \"%s\" not found", task.Name),
	}
}
