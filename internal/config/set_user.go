package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func write(con Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error finding the home directory: %w", err)
	}
	data, err := json.Marshal(con)
	if err != nil {
		return fmt.Errorf("error while marshalling the config: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("error while writing the config: %w", err)
	}
	return nil
}
