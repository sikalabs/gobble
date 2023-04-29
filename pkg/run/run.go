package run

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/play"
	"github.com/sikalabs/gobble/pkg/task"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

func RunFromFile(
	configFilePath string,
	dryRun bool,
	quietOutput bool,
	onlyTags []string,
) error {
	c, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	if c.Meta.SchemaVersion != 3 {
		return fmt.Errorf("unsupported schema version: %d", c.Meta.SchemaVersion)
	}

	c.AllPlays = []play.Play{}

	// Add Includes Before
	for _, includePlays := range c.IncludePlaysBefore {
		plays, err := getPlaysFromIncludePlays(includePlays)
		if err != nil {
			log.Fatalln(err)
		}
		c.AllPlays = append(c.AllPlays, plays...)
	}

	// Add Plays from Gobblefile
	c.AllPlays = append(c.AllPlays, c.Plays...)

	// Add Includes After
	for _, includePlays := range c.IncludePlaysAfter {
		plays, err := getPlaysFromIncludePlays(includePlays)
		if err != nil {
			log.Fatalln(err)
		}
		c.AllPlays = append(c.AllPlays, plays...)
	}

	lenPlays := lenPlays(c, onlyTags)
	playI := 0
	for _, play := range c.AllPlays {
		if len(onlyTags) > 0 {
			skip := true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
			if skip {
				continue
			}
		}
		playI++

		lenTasks := len(play.Tasks)
		taskI := 0
		for _, t := range play.Tasks {
			taskI++
			lenHosts := lenHosts(c, play)
			hostI := 0
			for globalHostName, globalHost := range c.Hosts {
				for _, host := range globalHost {
					if !slices.Contains(play.Hosts, globalHostName) {
						continue
					}
					hostI++

					if !quietOutput {
						fmt.Printf("+ play: %s (%d/%d)\n", play.Name, playI, lenPlays)
						fmt.Printf("  task: %s (%d/%d)\n", t.Name, taskI, lenTasks)
						fmt.Printf("  host: %s (%d/%d)\n", host.SSHTarget, hostI, lenHosts)
						if play.Sudo {
							fmt.Printf("  sudo: %t\n", play.Sudo)
						}
					}
					taskInput := libtask.TaskInput{
						SSHTarget:               host.SSHTarget,
						Config:                  c,
						NoStrictHostKeyChecking: c.Global.NoStrictHostKeyChecking,
						Sudo:                    play.Sudo,
						Vars:                    mergeMaps(c.Global.Vars, host.Vars),
						Dry:                     dryRun,
						Quiet:                   quietOutput,
					}
					out := task.Run(taskInput, t)
					if out.Error != nil {
						return out.Error
					}
					fmt.Println(``)
				}
			}
		}
	}

	return nil
}

func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	if m1 == nil {
		m1 = make(map[string]interface{})
	}
	deepCopyM1 := deepcopy.Copy(m1).(map[string]interface{})
	for k, v := range m2 {
		deepCopyM1[k] = v
	}
	return deepCopyM1
}

func readConfigFile(configFilePath string) (config.Config, error) {
	var buf []byte
	var err error
	c := config.Config{}

	if configFilePath == "-" {
		// Read from stdin
		buf, err = ioutil.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return c, err
		}
	} else {
		// Read from file
		buf, err = ioutil.ReadFile(configFilePath)
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

func lenPlays(c config.Config, onlyTags []string) int {
	length := 0
	for _, play := range c.AllPlays {
		if len(onlyTags) > 0 {
			skip := true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
			if skip {
				continue
			}
		}
		length++
	}

	return length
}

func lenHosts(c config.Config, play play.Play) int {
	length := 0
	for globalHostName, globalHost := range c.Hosts {
		for _, _ = range globalHost {
			if !slices.Contains(play.Hosts, globalHostName) {
				continue
			}
			length++
		}
	}
	return length
}

func getPlaysFromIncludePlays(includePlays config.InludePlays) ([]play.Play, error) {
	plays := []play.Play{}
	if strings.HasPrefix(includePlays.Source, "http://") ||
		strings.HasPrefix(includePlays.Source, "https://") {
		// Get from URL
		playsFromOneURL, err := getPlaysFromURL(includePlays.Source)
		if err != nil {
			return nil, err
		}
		plays = append(plays, playsFromOneURL...)
	} else {
		// Get from file
		playsFromOneFile, err := getPlaysFromFile(includePlays.Source)
		if err != nil {
			return nil, err
		}
		plays = append(plays, playsFromOneFile...)
	}
	return plays, nil
}

func getPlaysFromFile(filePath string) ([]play.Play, error) {
	var err error
	var buf []byte
	plays := []play.Play{}

	// Read from file
	buf, err = ioutil.ReadFile(filePath)
	if err != nil {
		return plays, err
	}

	_ = yaml.Unmarshal(buf, &plays)
	if err != nil {
		return plays, err
	}

	return plays, nil
}

func getPlaysFromURL(url string) ([]play.Play, error) {
	var err error
	var buf []byte
	plays := []play.Play{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	// Read from HTTP response
	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return plays, err
	}

	_ = yaml.Unmarshal(buf, &plays)
	if err != nil {
		return plays, err
	}

	return plays, nil
}
