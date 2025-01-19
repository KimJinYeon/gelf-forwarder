package internal

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	InboundPort  int    `yaml:"inbound_port"`
	OutboundHost string `yaml:"outbound_host"`
	OutboundPort int    `yaml:"outbound_port"`
}

// LoadConfig reads the configuration from config.yaml
func LoadConfig() (*Config, error) {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Println("Env CONFIG_PATH not found. Use default path: /config/config.yaml")
		// 기본 경로 지정
		configPath = "/config/config.yaml"
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
