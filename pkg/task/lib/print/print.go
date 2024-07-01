package print

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/printer"
	"github.com/sikalabs/gobble/pkg/utils"
	"os"
	text_template "text/template"

	"github.com/sikalabs/gobble/pkg/libtask"
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

	// Ensure the data is written and file is closed for reading
	err = tmpFile.Close()
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	// Read the content of the temporary file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	format := "OUTPUT:\n--------------------\n%s\n--------------------\n"
	printer.GlobalPrinter.Print(format, string(content))

	return libtask.TaskOutput{
		Error: nil,
	}
}
