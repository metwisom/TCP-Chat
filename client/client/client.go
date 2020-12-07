package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"../config"
)

func CreateConnect(config *config.Config) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err != nil {
		fmt.Println("Error connecting to server")
	}
	fmt.Println("Connection successfully established")
	return conn
}

func Writer(conn net.Conn) {
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
