package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ConfigChannel struct {
	Id                     string `yaml:"id"`
	Url                    string `yaml:"url"`
	StreamUrlPath          string `yaml:"streamUrlPath"`
	ContentType            string `yaml:"contentType"`
	ProgramUrl             string `yaml:"programUrl"`
	ProgramSelector        string `yaml:"programSelector"`
	ProgramStylesheet      string `yaml:"programStylesheet"`
	ProgramLocalStylesheet string `yaml:"programLocalStylesheet"`
	ProgramScrollTo        string `yaml:"programScrollTo"`
}

type Config struct {
	Port     int             `yaml:"port"`
	Channels []ConfigChannel `yaml:"channels"`
}

var config *Config

func readConfig() *Config {
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		panic(err)
	}

	return c
}

func GetConfig() *Config {
	if config == nil {
		config = readConfig()
	}

	return config
}
