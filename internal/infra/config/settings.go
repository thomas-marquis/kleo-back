package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Grpc     GrpcConfig     `yaml:"grpc"`
	Database DatabaseConfig `yaml:"database"`
}

type GrpcConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig(filePath string) (Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return Config{}, nil
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
