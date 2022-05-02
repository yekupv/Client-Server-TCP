package main

import (
	"bufio"
	"fmt"
	"net"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}
	conn, err := ln.Accept()

	if err != nil {
		panic(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			break
		}
		fmt.Print("Message Received:", string(message))

		conn.Write([]byte(message + "\n"))
	}
}
