package test

import (
	"github.com/BurntSushi/toml"
	"test/pkg/repository"
)

type Config struct {
	Port  string `toml:"port"`
	Mongo repository.MongoConfig
}

func NewConfig(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
