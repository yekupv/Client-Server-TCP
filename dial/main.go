package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your name: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("Hello, " + message)
	}
}
