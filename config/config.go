package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const _defaultConfigFile = "config/config.yaml"

type WebConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MysqlConfig struct {
	Dns string `yaml:"dns"`
}

type LogConfig struct {
	Output     int    `yaml:"output"`
	Formatter  int    `yaml:"formatter"`
	Level      int    `yaml:"level"`
	Folder     string `yaml:"folder"`
	ApiLogFile string `yaml:"api_log_file"`
	AppLogFile string `yaml:"app_log_file"`
}

type Config struct {
	Web   WebConfig   `yaml:"web"`
	Mysql MysqlConfig `yaml:"mysql"`
	Log   LogConfig   `yaml:"log"`
}

type ExternalConfig struct {
	GoogleUrl string `yaml:"google_url"`
}

type InternalConfig struct {
	Environment string `yaml:"environment"`
}

func NewConfig(configFile string) (*Config, error) {
	if configFile == "" {
		configFile = _defaultConfigFile
	}
	confData, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("loading config file failed: %v", err)
	}
	var config Config
	if err = yaml.Unmarshal(confData, &config); err != nil {
		return nil, fmt.Errorf("reading config failed: %v", err)
	}
	return &config, nil
}
