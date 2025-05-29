package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	var conf Config

	configPath, _ := getConfigFilePath()

	fileBytes, err := os.ReadFile(configPath)
	if err != nil {
		return conf, fmt.Errorf("Failed to read config file: %w", err)
	}

	err = json.Unmarshal(fileBytes, &conf)
	if err != nil {
		return conf, fmt.Errorf("Failed to unmarshal config JSON: %w", err)
	}

	return conf, nil
}
