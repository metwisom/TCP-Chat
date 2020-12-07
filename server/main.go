package main

import (
	"fmt"
	"os"

	"./config"
	"./server"
)

var cfg config.Config

func main() {
	cfg, err := config.LoadConfig("./config.json")
	if (err != nil && config.IsDefaultCrated(err)) || err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	server.CreateServer(cfg)
}
