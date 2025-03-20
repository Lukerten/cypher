package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Substitution struct {
		Keys string `yaml:"keys"`
	} `yaml:"substitution"`
	RSA struct {
		PublicKeyPath  string `yaml:"public_key_path"`
		PrivateKeyPath string `yaml:"private_key_path"`
	} `yaml:"rsa"`
	TripleDES struct {
		Key string `yaml:"key"`
	} `yaml:"3des"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
