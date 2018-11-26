package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config parameters from the configuration file
type Config struct {
	Store      string `json:"store"`
	Project    string `json:"project"`
	AWSProfile string `json:"aws_profile"`
	AWSRegion  string `json:"aws_region"`
}

// ReadConfig reads the team configuration
func ReadConfig() (*Config, error) {
	b, err := ioutil.ReadFile("team.json")
	if err != nil {
		return nil, errors.New("could not read team.json file")
	}

	c := &Config{}

	if err := json.Unmarshal(b, c); err != nil {
		return nil, errors.New("could not parse team.json file")
	}

	return c, nil
}
