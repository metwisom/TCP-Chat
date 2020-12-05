package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var config Config

func main() {
	config, err := LoadConfig("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		fmt.Printf("Server creation error: %s", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Client connection error: %s", err.Error())
		}
		fmt.Println("Client connected")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	go reader(conn)
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

func reader(conn net.Conn) {
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
