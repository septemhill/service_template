package config

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
)

const (
	jsonExt = ".json"
	ymlExt  = ".yml"
	yamlExt = ".yaml"
	xmlExt  = ".xml"
)

type MySQLConfig struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"db" yaml:"databas"`
}

type RWMySQLConfig struct {
	RDB MySQLConfig `json:"rdb" yaml:"rdb"`
	WDB MySQLConfig `json:"wdb" yaml:"wdb"`
}

type RedisConfig struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"db" yaml:"database"`
}

type KafkaConfig struct {
	Host string `json:"host" yaml:"host"`
}

type Infrastructure struct {
	MySQL   MySQLConfig   `json:"mysql" yaml:"mysql"`
	RWMySQL RWMySQLConfig `json:"rwmysql" yaml:"rwmysql"`
	Redis   RedisConfig   `json:"redis" yaml:"redis"`
	Kafka   KafkaConfig   `json:"kafka" yaml:"kafka"`
}

type External struct {
}

type Config struct {
	Infra    Infrastructure `json:"infra" yaml:"infar"`
	External External       `json:"json" yaml:"json"`
}

func isSupportExt(ext string) bool {
	return slices.Index([]string{jsonExt, yamlExt, ymlExt}, ext) >= 0
}

func NewConfig(path string) (*Config, error) {
	ext := filepath.Ext(path)

	if !isSupportExt(ext) {
		return nil, errors.New("unsupported file extension")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return decoderMap[ext](f)
}
