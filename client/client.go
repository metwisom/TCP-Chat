package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "10.180.91.44:8089")
	if err != nil {
		fmt.Println("Ошибка подключения")
	}
	fmt.Println("Соедниенение установлено")
	go reader(conn)
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println("second")
		message := []byte(input)
		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
	}
}

func reader(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err.Error() == "EOF" {
			fmt.Println("Собеседник отключился")
			break
		}
		fmt.Println("Собеседник:" + message)
	}
}
