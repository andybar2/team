package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configFile = "team.json"

// Config parameters from the configuration file
type Config struct {
	Store      string `json:"store"`
	Project    string `json:"project"`
	AWSProfile string `json:"aws_profile"`
	AWSRegion  string `json:"aws_region"`
}

// ReadConfig reads the team configuration
func ReadConfig() (*Config, error) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s file", configFile)
	}

	c := &Config{}

	if err := json.Unmarshal(b, c); err != nil {
		return nil, fmt.Errorf("could not parse %s file", configFile)
	}

	return c, nil
}
