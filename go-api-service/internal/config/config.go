package config

import (
	"os"
	"github.com/goccy/go-yaml"
)

type Config struct {
	PythonProcessorURL string   `yaml:"python_processor_url"`
	APIKeys            []string `yaml:"api_keys"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg Config
	// decoding the yaml file
	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	if envURL := os.Getenv("PYTHON_PROCESSOR_URL"); envURL != "" {
		cfg.PythonProcessorURL = envURL
	}
	if envKeys := os.Getenv("API_KEYS"); envKeys != "" {
		cfg.APIKeys = []string{envKeys}
	}
	return &cfg, nil
}
