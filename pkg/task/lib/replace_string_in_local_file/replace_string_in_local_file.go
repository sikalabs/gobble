package replace_string_in_local_file

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/template_utils"
)

type TaskReplaceStringInLocalFile struct {
	Path    string `yaml:"path"`
	Find    string `yaml:"find"`
	Replace string `yaml:"replace"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskReplaceStringInLocalFile,
) libtask.TaskOutput {
	vars := map[string]interface{}{
		"Config": taskInput.Config,
		"Vars":   taskInput.Vars,
	}

	// Render find string
	find, err := template_utils.RenderTemplateToString(
		taskParams.Find, "TaskReplaceStringInLocalFile.Find", vars,
	)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	// Render replace string
	replace, err := template_utils.RenderTemplateToString(
		taskParams.Replace, "TaskReplaceStringInLocalFile.Replace", vars,
	)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}

	err = replaceStringInFile(
		taskParams.Path,
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
	err = ioutil.WriteFile(filename, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
