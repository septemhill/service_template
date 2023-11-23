package config

import (
	"encoding/xml"
	"io"
)

func xmlDecoder(r io.Reader) (*Config, error) {
	var config Config

	if err := xml.NewDecoder(r).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
