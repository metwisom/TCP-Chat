package main

import (
	"fmt"
	"os"

	"./client"
	"./config"
)

func main() {
	cfg, err := config.LoadConfig("./config.json")
	if (err != nil && config.IsDefaultCrated(err)) || err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn := client.CreateConnect(cfg)
	go client.Reader(conn)
	client.Writer(conn)
}
