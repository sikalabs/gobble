package print

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/utils"
	"os"
	text_template "text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type Task struct {
	libtask.BaseTask
	Template string `yaml:"template"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	tmpl, err := text_template.New("template").Parse(t.Template)
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
	defer os.Remove(tmpFile.Name())
	err = tmpl.Execute(tmpFile, map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   utils.MergeMaps(taskInput.Vars, host.Vars),
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	if taskInput.Quiet {
		fmt.Println("cat <<EOF")
	}

	fmt.Println("OUTPUT:")
	err = exec_utils.RawExecStdout("cat", tmpFile.Name())
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	if taskInput.Quiet {
		fmt.Println("EOF")
	}

	return libtask.TaskOutput{
		Error: nil,
	}
}
