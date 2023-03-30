package libtask

type TaskInput struct {
	SSHTarget               string
	NoStrictHostKeyChecking bool
	Sudo                    bool
	Config                  interface{}
	Vars                    map[string]interface{}
	Dry                     bool
	Quiet                   bool
}

type TaskOutput struct {
	Error error
}
