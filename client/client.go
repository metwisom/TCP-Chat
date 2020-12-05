package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var config Config

func main() {
	LoadConfig(&config)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err != nil {
		fmt.Println("Error connecting to server")
	}
	fmt.Println("Connection successfully established")
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
