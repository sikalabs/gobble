package replace_string_in_local_file

import (
	"io/ioutil"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
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
	err := replaceStringInFile(taskParams.Path, taskParams.Find, taskParams.Replace)
	return libtask.TaskOutput{
		Error: err,
	}
}

func replaceStringInFile(filename, searchStr, replaceStr string) error {
	// read the entire file into memory
	content, err := ioutil.ReadFile(filename)
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
