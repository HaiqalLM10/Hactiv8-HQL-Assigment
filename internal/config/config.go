package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Token string `yaml:"token"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var Cfg *Config

func Load(filename string) (err error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	// process umarshal to struct config
	err = yaml.Unmarshal(f, &Cfg)
	return
}
