package play

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetPlaysFromIncludePlays(includePlays InludePlays) ([]Play, error) {
	var plays []Play
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

func getPlaysFromFile(filePath string) ([]Play, error) {
	var err error
	var buf []byte
	var plays []Play

	// Read from file
	buf, err = os.ReadFile(filePath)
	if err != nil {
		return plays, err
	}

	_ = yaml.Unmarshal(buf, &plays)
	if err != nil {
		return plays, err
	}

	return plays, nil
}

func getPlaysFromURL(url string) ([]Play, error) {
	var err error
	var buf []byte
	var plays []Play

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	// Read from HTTP response
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		return plays, err
	}

	_ = yaml.Unmarshal(buf, &plays)
	if err != nil {
		return plays, err
	}

	return plays, nil
}
