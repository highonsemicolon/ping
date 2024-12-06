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

	Auth0 struct {
		Domain   string `yaml:"domain"`
		Audience string `yaml:"audience"`
	} `yaml:"auth0"`

	MySQL struct {
		DSN        string `yaml:"dsn"`
		CACertPath string `yaml:"ca_cert_path"`
	} `yaml:"mysql"`
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
