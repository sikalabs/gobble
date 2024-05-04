package template

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/utils"
	"os"
	text_template "text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/chmod"
	"github.com/sikalabs/gobble/pkg/task/lib/cp"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type Task struct {
	libtask.BaseTask
	Path      string      `yaml:"path"`
	Template  string      `yaml:"template"`
	ExtraData interface{} `yaml:"extra_data"`
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
	err = tmpl.Execute(tmpFile, map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   utils.MergeMaps(taskInput.Vars, host.Vars),
		"Extra":  t.ExtraData,
	})
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	if taskInput.Dry {
		fmt.Println("cat > " + tmpFile.Name() + " <<EOF")
		exec_utils.RawExecStdout("cat", tmpFile.Name())
		fmt.Println("EOF")
	}

	//cp Task
	cpTask := cp.Task{
		BaseTask: libtask.BaseTask{
			Name: t.Name,
		},
		LocalSrc:  tmpFile.Name(),
		RemoteDst: t.Path,
	}

	out := cpTask.Run(taskInput, host)
	if out.Error != nil {
		return libtask.TaskOutput{
			Error: out.Error,
		}
	}
	chmod := chmod.Task{Path: t.Path, Perm: "644"}
	out = chmod.Run(taskInput, host)
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
