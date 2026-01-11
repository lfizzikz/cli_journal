package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveConfig(cfg Config) error {
	err := cfg.Validate()
	if err != nil {
		return err
	}

	cfgJSON, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFile, err := GetConfigFile()
	if err != nil {
		return err
	}

	dir := filepath.Dir(configFile)
	os.MkdirAll(dir, 0755)
	err = os.WriteFile(configFile, cfgJSON, 0600)
	if err != nil {
		return err
	}

	return nil
}
