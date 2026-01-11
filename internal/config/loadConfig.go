package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig() (Config, error) {
	configFile, err := GetConfigFile()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory: %v", err)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error reading json file: %v", err)
	}

	return cfg, nil
}

func GetConfigFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory %v", err)
	}
	configFile := filepath.Join(homeDir, ".config", "cli_obsidian", "config.json")
	return configFile, nil
}
