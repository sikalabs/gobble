package yaml_utils

import (
	"fmt"
	"os"

	"github.com/sikalabs/gobble/pkg/logger"

	"gopkg.in/yaml.v3"
)

func MergeYAMLs(paths []string) {
	var merged map[interface{}]interface{}

	data := []byte{}

	for _, path := range paths {
		file, err := os.ReadFile(path)
		if err != nil {
			logger.Log.Fatalf("Failed to read %s: %v", path, err)
		}
		data = append(data, file...)
	}

	err := yaml.Unmarshal(data, &merged)
	if err != nil {
		logger.Log.Fatalf("Failed to merge YAML files: %v", err)
	}

	// Print the merged YAML
	out, err := yaml.Marshal(merged)
	if err != nil {
		logger.Log.Fatalf("Failed to marshal merged YAML: %v", err)
	}
	fmt.Println(string(out))
}
