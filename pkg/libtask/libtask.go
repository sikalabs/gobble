package libtask

type TaskInput struct {
	SSHTarget string
	Config    interface{}
	Vars      map[string]string
	Dry       bool
}

type TaskOutput struct {
	Error error
}
