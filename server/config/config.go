package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	LoadError          = "Config file loading error: %s\n"
	DefaultCreateError = "Error creating default config: %s\n"
	CreateDefault      = "Default config created, configure it\n"
	CloseError         = "Error when closing file: %s\n"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

func (c *Config) setDefault() {
	c.Server.Port = 8089
}

func LoadConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := createDefaultConfig(path)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(CreateDefault)
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf(LoadError, err.Error()))
	}
	bConfig, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	return decodeConfig(bConfig)
}

func IsDefaultCrated(err error) bool {
	return err.Error() == CreateDefault
}

func readConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LoadError, err.Error()))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf(CloseError, err.Error())
		}
	}()
	bConfig, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LoadError, err.Error()))
	}
	return bConfig, nil
}

func decodeConfig(bConfig []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(bConfig, &config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(LoadError, err.Error()))
	}
	return &config, nil
}

func createDefaultConfig(path string) error {
	var config Config
	config.setDefault()
	bConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf(DefaultCreateError, err.Error()))
	}
	err = ioutil.WriteFile(path, bConfig, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf(DefaultCreateError, err.Error()))
	}
	return nil
}
