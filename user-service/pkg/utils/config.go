package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	MySQL MySQLConfig `yaml:"mysql"`

	Auth0 Auth0 `yaml:"auth0"`
}

type Auth0 struct {
	Domain   string `yaml:"domain"`
	Audience string `yaml:"audience"`
}

type MySQLConfig struct {
	DSN        string `yaml:"dsn"`
	CACertPath string `yaml:"ca_cert_path"`
}

func LoadConfig(filepath string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
