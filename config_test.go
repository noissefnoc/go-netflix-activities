package main

import (
	"testing"
	"os"
)

func TestReadNormalWEnvCase(t *testing.T) {
	os.Setenv("NF_EMAIL", "env@example.com")
	os.Setenv("NF_PASSWORD", "envenv")

	config := &Config{}

	err := config.Read("./test/resources/config/test.toml")

	if err != nil {
		t.Fatalf("failed pass environmental variables %#v", err)
	}

	if config.Auth.Email != "env@example.com" {
		t.Errorf("failed to read config value 'Email' %#v", err)
	}

	if config.Auth.Password != "envenv" {
		t.Errorf("failed to read config value 'Password' %#v", err)
	}

	// ensure clear NF_EMAIL and NF_PASSWORD
	os.Clearenv()
}

func TestReadNormalWoEnvCase(t *testing.T) {
	config := &Config{}

	err := config.Read("./test/resources/config/test.toml")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if config.Auth.Email != "hoge@examples.com" {
		t.Errorf("failed to read config value 'Email' %#v", err)
	}

	if config.Auth.Password != "hogehoge" {
		t.Errorf("failed to read config value 'Password' %#v", err)
	}
}

func TestReadRaiseErrorCase(t *testing.T) {
	config := &Config{}

	err := config.Read("./test/resources/does_not_exist.toml")

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
}
