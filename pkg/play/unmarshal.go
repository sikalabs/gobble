package play

import (
	"fmt"
	"github.com/sikalabs/gobble/pkg/task"
	"gopkg.in/yaml.v3"
)

func (p *Play) UnmarshalYAML(node *yaml.Node) error {
	type alias Play // Use an alias to avoid recursion
	var tempPlay alias

	if err := node.Decode(&tempPlay); err != nil {
		return err
	}

	// Set the simple fields
	p.Name = tempPlay.Name
	p.Sudo = tempPlay.Sudo
	p.Tags = tempPlay.Tags
	p.Hosts = tempPlay.Hosts

	// Now handle the tasks array
	for i := 0; i < len(node.Content); i++ {
		if node.Content[i].Kind == yaml.ScalarNode && node.Content[i].Value == "tasks" {
			tasksNode := node.Content[i+1]
			for _, tNode := range tasksNode.Content {
				if tNode.Kind == yaml.MappingNode {
					var taskName string
					var taskType string
					var taskParamsNode *yaml.Node

					for j := 0; j < len(tNode.Content); j += 2 {
						key := tNode.Content[j].Value
						valNode := tNode.Content[j+1].Value
						if key == "name" {
							taskName = valNode
						} else if key != "name" {
							taskType = key
							taskParamsNode = tNode.Content[j+1]
						}
					}

					if taskType != "" {
						constructor, exists := task.Registry[taskType]
						if exists {
							newTask := constructor()
							newTask.SetName(taskName)
							// Decode parameters into the new task object
							if err := taskParamsNode.Decode(newTask); err != nil {
								return fmt.Errorf("failed to decode %s task: %v", taskType, err)
							}
							p.Tasks = append(p.Tasks, newTask)
						} else {
							return fmt.Errorf("unknown task type: %s", taskType)
						}
					}
				}
			}
			break // We found the tasks array, no need to continue
		}
	}

	return nil
}
