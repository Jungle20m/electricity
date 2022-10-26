package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type WebConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MysqlConfig struct {
	Dns string `yaml:"dns"`
}

type LogConfig struct {
	Output    string `yaml:"output"`
	Formatter string `yaml:"formatter"`
	Level     string `yaml:"level"`
}

type ElectricityConfig struct {
	Web   WebConfig   `yaml:"web"`
	Mysql MysqlConfig `yaml:"dns"`
	Log   LogConfig   `yaml:"log"`
}

var ServiceConfig ElectricityConfig

func LoadConfig(filename string) error {
	confData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("loading config file failed: %v", err)
	}
	if err = yaml.Unmarshal(confData, &ServiceConfig); err != nil {
		return fmt.Errorf("reading config failed: %v", err)
	}
	return nil
}
