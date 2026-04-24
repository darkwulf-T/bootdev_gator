package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error finding the home directory: %w", err)
	}
	f, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error while trying to opne the file: %w", err)
	}
	defer f.Close()

	var config Config
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("error decoding the config file: %w", err)
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homePath, configFileName)
	return configPath, nil
}
