package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = "config.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(userName string) error {
	if userName == "" {
		return errors.New("you need to set a user name")
	}

	c.CurrentUserName = userName

	err := write(c)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var config Config

	path, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(home, ".config", "gogator")
	err = os.MkdirAll(configPath, 0744)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(configPath, configFileName)

	return filePath, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0640)
	if err != nil {
		return err
	}

	return nil
}
