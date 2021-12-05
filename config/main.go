package config

import (
	"github.com/goccy/go-yaml"
	"github.com/stdcurse/pm/output"
	"io/ioutil"
)

type Config struct {
	Repos []struct {
		Url  string
		Path string
	}
	Portdir  string
	Database string
	Env      map[string]string
}

func NewConfig(path string) *Config {
	content, err := ioutil.ReadFile(path)
	output.Check(err, "Something went wrong with config file", true)

	var c Config
	output.Check(yaml.Unmarshal(content, &c), "Something went wrong with config file content", true)

	return &c
}
