package libtask

type TaskInput struct {
	SSHTarget               string
	SSHPort                 int
	SSHPassword             string
	SSHOptions              []string
	NoStrictHostKeyChecking bool
	Sudo                    bool
	SudoPassword            string
	Config                  interface{}
	Vars                    map[string]interface{}
	Dry                     bool
	Quiet                   bool
}

type TaskOutput struct {
	Error error
}

// BaseTask provides common fields for all tasks.
type BaseTask struct {
	Name string
}

// SetName sets the name of the task.
func (bt *BaseTask) SetName(name string) {
	bt.Name = name
}

// GetName returns the name of the task.
func (bt *BaseTask) GetName() string {
	return bt.Name
}
