package config

import (
	"encoding/json"
	"io"
)

func jsonDecoder(r io.Reader) (*Config, error) {
	var config Config

	if err := json.NewDecoder(r).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
