package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find user home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)

	return configPath, nil
}

func write(conf Config) error {
	configPath, _ := getConfigFilePath()

	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
