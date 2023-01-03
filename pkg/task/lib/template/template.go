package template

import (
	"os"
	text_template "text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/chmod"
	"github.com/sikalabs/gobble/pkg/task/lib/cp"
)

type TaskTemplete struct {
	Path      string      `yaml:"path"`
	Template  string      `yaml:"template"`
	ExtraData interface{} `yaml:"extra_data"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskTemplete,
) libtask.TaskOutput {
	tmpl, err := text_template.New("template").Parse(taskParams.Template)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	tmpFile, err := os.CreateTemp("", "template")
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	err = tmpl.Execute(tmpFile, map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   taskInput.Vars,
		"Extra":  taskParams.ExtraData,
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	out := cp.Run(taskInput, cp.TaskCp{
		LocalSrc:  tmpFile.Name(),
		RemoteDst: taskParams.Path,
	})
	if out.Error != nil {
		return libtask.TaskOutput{
			Error: out.Error,
		}
	}
	out = chmod.Run(taskInput, chmod.TaskChmod{
		Path: taskParams.Path,
		Perm: "644",
	})
	if out.Error != nil {
		return libtask.TaskOutput{
			Error: out.Error,
		}
	}
	defer os.Remove(tmpFile.Name())

	return libtask.TaskOutput{
		Error: err,
	}
}
