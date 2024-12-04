// Задача:
// Создайте слайс целых чисел и заполните его числами от 1 до 10.
// Используя цикл, пройдите по слайсу и увеличьте каждое значение на 5, используя указатель.
// Выведите измененный слайс.

package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 10)
	
	for i := 0; i < 10; i++ {
		slice[i] = i + 1
	}

	for i := range slice {
		ptr := &slice[i]
		*ptr += 5
	}

	fmt.Println("New slice:", slice)

}