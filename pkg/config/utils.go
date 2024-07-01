package config

import (
	"github.com/sikalabs/gobble/pkg/host"
)

func LenTargets(targets host.Targets) int {
	length := 0
	for _, h := range targets {
		for _, _ = range h {
			length++
		}
	}
	return length
}
