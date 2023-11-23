package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

func yamlDecoder(r io.Reader) (*Config, error) {
	var config Config

	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
