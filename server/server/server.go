package server

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"../config"
)

func CreateServer(cfg *config.Config) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		fmt.Printf("Server creation error: %s", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Client connection error: %s", err.Error())
		}
		fmt.Println("Client connected")
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	go Reader(conn)
	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		message := []byte(input)
		length, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Send error: %s\nBytes written: %d", err.Error(), length)
			break
		}
	}
}

func Reader(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Server connection interrupted")
			fmt.Println(err.Error())
			break
		}
		fmt.Println("Received: " + message)
	}
}
