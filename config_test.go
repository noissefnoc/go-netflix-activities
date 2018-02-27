package main

import "testing"

func TestReadNormalCase(t *testing.T) {
	config := &Config{}

	err := config.Read("./test/resources/test.toml")

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
