package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigSet Configuration

type Configuration struct {
	Name       string            `yaml:"name"`
	Properties map[string]string `yaml:"properties"`
}

func InitializeConfigurations() {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, &ConfigSet)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
