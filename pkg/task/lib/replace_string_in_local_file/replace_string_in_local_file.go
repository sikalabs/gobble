package replace_string_in_local_file

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/utils"
	"os"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/template_utils"
)

type Task struct {
	libtask.BaseTask
	Path    string `yaml:"path"`
	Find    string `yaml:"find"`
	Replace string `yaml:"replace"`
}

func (t *Task) Run(taskInput libtask.TaskInput, host *host.Host) libtask.TaskOutput {
	vars := map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   utils.MergeMaps(taskInput.Vars, host.Vars),
	}

	// Render find string
	find, err := template_utils.RenderTemplateToString(
		t.Find, "TaskReplaceStringInLocalFile.Find", vars,
	)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	// Render replace string
	replace, err := template_utils.RenderTemplateToString(
		t.Replace, "TaskReplaceStringInLocalFile.Replace", vars,
	)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	err = replaceStringInFile(
		t.Path,
		find,
		replace,
	)
	return libtask.TaskOutput{
		Error: err,
	}
}

func replaceStringInFile(filename, searchStr, replaceStr string) error {
	// read the entire file into memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// replace the string
	newContent := strings.ReplaceAll(string(content), searchStr, replaceStr)

	// write the updated content back to the file
	err = os.WriteFile(filename, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
