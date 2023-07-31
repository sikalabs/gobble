package echo

import (
	"bytes"
	"text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskEcho struct {
	Message string `yaml:"message"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskEcho,
) libtask.TaskOutput {

	tmpl, err := template.New("cmd").Parse(taskParams.Message)
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
		"echo "+buf.String(),
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
