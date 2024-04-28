package config

import (
	"github.com/sikalabs/gobble/pkg/host"
	"github.com/sikalabs/gobble/pkg/play"
	"golang.org/x/exp/slices"
)

func LenPlays(c *Config, onlyTags []string, skipTags []string) int {
	length := 0
	for _, play := range c.AllPlays {
		skip := false
		for _, tag := range skipTags {
			if slices.Contains(play.Tags, tag) {
				skip = true
			}
		}
		if skip {
			continue
		}
		if len(onlyTags) > 0 {
			skip = true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
		}
		if skip {
			continue
		}
		length++
	}

	return length
}

func LenHosts(c *Config, play play.Play) int {
	length := 0
	for globalHostName, globalHost := range c.AllHosts {
		for _, _ = range globalHost {
			if !slices.Contains(play.Hosts, globalHostName) {
				continue
			}
			length++
		}
	}
	return length
}

func LenTargets(targets host.Targets) int {
	length := 0
	for _, h := range targets {
		for _, _ = range h {
			length++
		}
	}
	return length
}
