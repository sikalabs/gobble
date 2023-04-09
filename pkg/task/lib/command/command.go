package command

import (
	"bytes"
	"text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskCommand struct {
	Cmd string `yaml:"cmd"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskCommand,
) libtask.TaskOutput {

	tmpl, err := template.New("cmd").Parse(taskParams.Cmd)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   taskInput.Vars,
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	err = exec_utils.SSH(
		taskInput,
		buf.String(),
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
