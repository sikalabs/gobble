package config

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func ReadConfigFile(configFilePath string) (*Config, error) {
	var buf []byte
	var err error
	var c *Config

	if configFilePath == "-" {
		// Read from stdin
		buf, err = io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return c, err
		}
	} else {
		// Read from file
		buf, err = os.ReadFile(configFilePath)
		if err != nil {
			return c, err
		}
	}

	_ = yaml.Unmarshal(buf, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
