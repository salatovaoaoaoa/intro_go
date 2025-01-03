// Разобраться и пофиксить:
// Если отправляем пустое сообщение, то клиент отключается,
// ничего не получаем и мы просто висим

// Ответ
// Если у нас длина сообщения = 0 (lenRequest = 0), то здесь: lenRequest, err := conn.Read(buffer) у нас будет 0 байт,
// => клиент ожидает данные, которых нет, поэтому виснет (у нас продолжает выполняться цикл,
// а клиентское соединение не закрывается)

// Пофикшенный код
// нужно добавить проверку если lenRequest = 0, ниже исправленный код:
// ??
// Раскомментила??:
// Если клиент прислал пустое сообщение
// if lenRequest == 0 {
// 	fmt.Println("Client sent empty message, closing connection")
// 	break
// }

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

	buffer := make([]byte, 1024) // Буфер для чтения данных

	for {
		// Читаем данные от клиента
		lenRequest, err := conn.Read(buffer)
		if err != nil {
			// Проверяем, если клиент закрыл соединение
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected:", conn.RemoteAddr().String())
				break
			}
			// Логируем другие ошибки
			fmt.Println("Error reading from client:", err.Error())
			break
		}

		// Если клиент прислал пустое сообщение (lenRequest == 0)
		if lenRequest == 0 {
			fmt.Println("Client sent empty message, closing connection...")
			break // прерываю
		}

		// Выводим полученные данные
		fmt.Printf("Bytes received: %d, Data: %s\n", lenRequest, string(buffer[:lenRequest]))

		// Отправляем ответ клиенту
		_, err = conn.Write([]byte("Hello from Server\n"))
		if err != nil {
			fmt.Println("Error sending data to client:", err)
			break
		}
	}
}
