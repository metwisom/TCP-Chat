package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

func (c *Config) set_default() {
	c.Server.Port = 8089
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := os.Stat(path); err == nil {
		cfg, err := os.Open(path)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Config file loading error: %s\n", err.Error()))
		}
		decoder := json.NewDecoder(cfg)
		err = decoder.Decode(&config)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Config file loading error: %s\n", err.Error()))
		}
	} else if os.IsNotExist(err) {
		config.set_default()
		err := createDefaultConfig(path, &config)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return nil, errors.New("Default config created")
	} else {
		return nil, errors.New(fmt.Sprintf("Config file loading error: %s\n", err.Error()))
	}
	return &config, nil
}

func createDefaultConfig(path string, config *Config) error {
	config_string, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating default config: %s\n", err.Error()))
	}
	err = ioutil.WriteFile(path, config_string, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating default config: %s\n", err.Error()))
	}
	return nil
}
