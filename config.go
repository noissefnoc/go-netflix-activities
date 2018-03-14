package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"os"
)

// Config is struct for this program
type Config struct {
	Auth AuthConfig
}

// AuthConfig is struct for Netflix Authentication
type AuthConfig struct {
	Email string
	Password string
}

// Read is API to read configurations from environmental variable or toml file
func (c *Config) Read(path string) (err error) {
	email := os.Getenv("NF_EMAIL")
	password := os.Getenv("NF_PASSWORD")

	if len(email) != 0 && len(password) != 0 {
		c.Auth.Email = email
		c.Auth.Password = password
	} else {
		_, err = toml.DecodeFile(path, c)

		if err != nil {
			return fmt.Errorf("faild to read config:%v", err)
		}
	}

	return nil
}
