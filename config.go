package test

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Port string `toml:"port"`
}

func NewConfig(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
