package libtask

type TaskInput struct {
	SSHTarget               string
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
