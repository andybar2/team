package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config parameters
type Config struct {
	Store      string `json:"store"`
	Project    string `json:"project"`
	AWSProfile string `json:"aws_profile"`
	AWSRegion  string `json:"aws_region"`
}

// ReadConfig reads the configuration from the given path
func ReadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read config file")
	}

	c := &Config{}

	if err := json.Unmarshal(b, c); err != nil {
		return nil, errors.New("could not parse config file")
	}

	return c, nil
}
