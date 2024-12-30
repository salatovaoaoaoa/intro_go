// Реализуйте простую модель многопользовательского чата, где несколько пользователей могут отправлять сообщения
// в общий канал. Каждое сообщение должно содержать имя отправителя и текст сообщения.
// Создайте несколько горутин для имитации пользователей, которые отправляют сообщения.

// Примерный вывод:

// [User3]: Message 1 from User3
// [User1]: Message 1 from User1
// [User2]: Message 1 from User2
// [User2]: Message 2 from User2
// [User1]: Message 2 from User1
// [User3]: Message 2 from User3
// [User3]: Message 3 from User3
// [User2]: Message 3 from User2
// [User1]: Message 3 from User1
// [User2]: Message 4 from User2
// [User3]: Message 4 from User3
// [User1]: Message 4 from User1
// [User1]: Message 5 from User1
// [User3]: Message 5 from User3
// [User2]: Message 5 from User2

package main

import (
	"fmt"
	"sync"
)

func user(ch chan string, userName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 5; j++ {
		ch <- fmt.Sprintf("[%s]: Message %d from %s", userName, j, userName)
	}
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go user(ch, fmt.Sprintf("User%d", i+1), &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
