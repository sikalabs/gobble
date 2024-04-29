package task

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/authorized_key"
	"github.com/sikalabs/gobble/pkg/task/lib/command"
)

// Task is the interface that all tasks need to implement.
type Task interface {
	Run(input libtask.TaskInput, host *host.Host) libtask.TaskOutput
	GetName() string
	SetName(name string)
}

// Constructor is a function signature for task constructors.
type Constructor func() Task

// Registry maps a task type name to its constructor.
var Registry = make(map[string]Constructor)

// RegisterTask adds a new task type to the registry.
func RegisterTask(name string, constructor Constructor) {
	if _, exists := Registry[name]; exists {
		panic(fmt.Sprintf("task type %q is already registered", name))
	}
	Registry[name] = constructor
}

// NewTask creates a new task by type name.
func NewTask(typeName string) (Task, error) {
	constructor, ok := Registry[typeName]
	if !ok {
		return nil, fmt.Errorf("no task registered with type %q", typeName)
	}
	return constructor(), nil
}

func init() {
	RegisterTask("command", func() Task { return &command.Task{} })
	RegisterTask("authorized_key", func() Task { return &authorized_key.Task{} })
}
