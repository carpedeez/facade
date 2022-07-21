package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	FacadeConfig `yaml:"facade"`
	DBConfig     `yaml:"db"`
	LogConfig    `yaml:"log"`
}

type FacadeConfig struct {
	Port string `yaml:"port"`
}
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type LogConfig struct{}

func GetConfig() (Config, error) {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %v", err)
	}

	c := Config{FacadeConfig{}, DBConfig{}, LogConfig{}}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config into struct: %v", err)
	}
	return c, nil
}
