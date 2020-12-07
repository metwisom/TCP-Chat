package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	LOAD_ERROR           = "Config file loading error: %s\n"
	DEFAULT_CREATE_ERROR = "Error creating default config: %s\n"
	DEFAULT_CREATED      = "Default config created, configure it"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
}

func (c *Config) set_default() {
	c.Server.Host = "localhost"
	c.Server.Port = 8089
}

func LoadConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := createDefaultConfig(path)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(DEFAULT_CREATED)
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf(LOAD_ERROR, err.Error()))
	}
	bConfig, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	return decodeConfig(bConfig)
}

func IsDefaultCrated(err error) bool {
	return err.Error() == DEFAULT_CREATED
}

func readConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LOAD_ERROR, err.Error()))
	}
	defer file.Close()
	bConfig, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LOAD_ERROR, err.Error()))
	}
	return bConfig, nil
}

func decodeConfig(bConfig []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(bConfig, &config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LOAD_ERROR, err.Error()))
	}
	return &config, nil
}

func createDefaultConfig(path string) error {
	var config Config
	config.set_default()
	bConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf(DEFAULT_CREATE_ERROR, err.Error()))
	}
	err = ioutil.WriteFile(path, bConfig, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf(DEFAULT_CREATE_ERROR, err.Error()))
	}
	return nil
}
