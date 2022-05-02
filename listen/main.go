package main

import (
	"bufio"
	"fmt"
	"net"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", ":8081")

	if err != nil { 
		panic(err)
	}
	// Открываем порт
	conn, err := ln.Accept()

	if err != nil { 
		panic(err)
	}

	// Запускаем цикл
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, err := bufio.NewReader(conn).ReadString('\n')
		
		if err != nil {
			break
		}
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", string(message))

		// Отправить новую строку обратно клиенту
		conn.Write([]byte(message + "\n"))
	}
}
