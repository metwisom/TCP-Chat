package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
}

func LoadConfig(config *Config) {
	cfg, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	decoder := json.NewDecoder(cfg)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("Config file loading error: %s\n", err.Error())
		os.Exit(3)
	}
}
