package libtask

type TaskInput struct {
	SSHTarget string
	Config    interface{}
	Vars      map[string]interface{}
	Dry       bool
}

type TaskOutput struct {
	Error error
}
