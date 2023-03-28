package print

import (
	"fmt"
	"os"
	text_template "text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskPrint struct {
	Template string `yaml:"template"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskPrint,
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
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	fmt.Println("OUTPUT:")
	err = exec_utils.RawExecStdout("cat", tmpFile.Name())
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	defer os.Remove(tmpFile.Name())

	return libtask.TaskOutput{
		Error: err,
	}
}
