// Напишите программу, которая принимает строки через канал и подсчитывает количество слов в каждой строке.
// Используйте несколько горутин для обработки строк и один канал для передачи результатов.

// Условно, на вход строка
// "Всем привет!
// Следующая лекция в среду
// Увидимся на лекции!
// Результат
// Word count: 2
// Word count: 4
// Word count: 3

package main

import (
	"fmt"
	"strings"
	"sync"
)

type InputMessage struct {
	Index int
	Line  string
}

type OutputMessage struct {
	Index     int
	WordCount int
}

func wordCounter(indexedIn <-chan InputMessage, out chan<- OutputMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range indexedIn {
		wordCount := len(strings.Fields(item.Line))
		out <- OutputMessage{Index: item.Index, WordCount: wordCount}
	}
}

func main() {
	indexedIn := make(chan InputMessage)
	out := make(chan OutputMessage)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go wordCounter(indexedIn, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	go func() {
		lines := []string{
			"Всем привет!",
			"Следующая лекция в среду",
			"Увидимся на лекции!",
		}
		for i, line := range lines {
			indexedIn <- InputMessage{Index: i, Line: line}
		}
		close(indexedIn)
	}()

	// использую мапу для упорядочивания
	results := make(map[int]int)
	for result := range out {
		results[result.Index] = result.WordCount
	}

	for i := 0; i < len(results); i++ {
		fmt.Printf("Word count: %d\n", results[i])
	}
}
