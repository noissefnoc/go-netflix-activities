package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	Email string
	Password string
}

func (c *Config) Read(path string) (err error) {
	_, err = toml.DecodeFile(path, c)

	if err != nil {
		return fmt.Errorf("faild to read config:%v", err)
	}

	return nil
}
