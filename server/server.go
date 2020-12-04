package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, _ := net.Listen("tcp", ":8089")
	for {
		conn, _ := ln.Accept()
		fmt.Println("Собеседник подключился")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader(conn)
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		message := []byte(input)
		conn.Write(message)
	}
}

func reader(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Собеседник:" + message)
	}
}
