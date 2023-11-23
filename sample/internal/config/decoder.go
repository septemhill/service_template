package config

import "io"

type decoderFunc func(io.Reader) (*Config, error)

var decoderMap = map[string]decoderFunc{
	jsonExt: jsonDecoder,
	ymlExt:  yamlDecoder,
	yamlExt: yamlDecoder,
	xmlExt:  xmlDecoder,
}
