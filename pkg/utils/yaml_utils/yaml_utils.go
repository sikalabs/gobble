package yaml_utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func MergeYAMLs(paths []string) {
	var merged map[interface{}]interface{}

	data := []byte{}

	for _, path := range paths {
		file, err := os.ReadFile(string(path))
		if err != nil {
			log.Fatalf("Failed to read %s: %v", string(path), err)
		}
		data = append(data, file...)
	}

	err := yaml.Unmarshal(data, &merged)
	if err != nil {
		log.Fatalf("Failed to merge YAML files: %v", err)
	}

	// Print the merged YAML
	out, err := yaml.Marshal(merged)
	if err != nil {
		log.Fatalf("Failed to marshal merged YAML: %v", err)
	}
	fmt.Println(string(out))
}
