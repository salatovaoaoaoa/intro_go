// Если мы подключаемся к этому порту, то он выводит ошибку End Of file (EOF)
// не можем прочитать, клиент отключился
// Здадание: почему код возвращает такую ошибку?
// пофиксить

// Ответ

// Ошибка End Of file (EOF) возникает, когда клиент отправляет данные и закрывает соединение или не отправляет данные,
// а сервер продолжает пытаться прочитать (в handleConnection),
// но уже из закрытого соединения (метод conn.Read(buffer) возвращает ошибку с сообщением EOF).
// Нужно ее корректно обработать: просто проверить на наличие этой ошибки и прервать попытки соединения

// Пофикшенный код

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Открываем порт для прослушивания
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8081...")

	for {
		// Ждем подключений от клиентов
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			os.Exit(1)
		}

		// Обрабатываем подключение в отдельной горутине
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	for {
		// Чтение данных от клиента в виде байтов
		buffer := make([]byte, 1024) // Создаем буфер для чтения данных
		lenRequest, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
			// Клиент закрыл соединение, завершение
				fmt.Println("Client disconnected:", conn.RemoteAddr().String())
				break
			}
			// Лог других ошибок
			fmt.Println("Error reading from client:", err.Error())
			break
		}

		fmt.Println("Bytes received:", lenRequest)
		// Jтправляем обратно  сообщение
		_, err = conn.Write([]byte{72, 101, 108, 108, 111, 10})
		if err != nil {
			fmt.Println("Error sending data to client:", err)
			return

		}
	}
}
