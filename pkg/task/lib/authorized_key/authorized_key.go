package authorized_key

import (
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskAuthorizedKey struct {
	Key string `yaml:"key"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskAuthorizedKey,
) libtask.TaskOutput {
	err := exec_utils.SSH(
		taskInput,
		"mkdir -p ~/.ssh",
	)
	if err != nil {
		return libtask.TaskOutput{
			Error: err,
		}
	}
	err = exec_utils.SSH(
		taskInput,
		`grep -qxF "`+taskParams.Key+`" ~/.ssh/authorized_keys || echo "`+taskParams.Key+`" >> ~/.ssh/authorized_keys`,
	)
	return libtask.TaskOutput{
		Error: err,
	}
}
