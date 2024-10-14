package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(username string) error {
	config.CurrentUserName = username

	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		return nil
	}

	encoder := json.NewEncoder(f)
	encoder.Encode(config)
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home path: %w", err)
	}
	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json: %w", err)
	}
	return config, nil
}
