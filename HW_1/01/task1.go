// Задача:
// Дан слайс
// Необходимо заменить четные числа на 1 и посчитать сумму чисел в новом слайсе
// numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}

package main

import (
	"fmt"
)

func main() {
	
	numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}

	sum := 0
	for i := range numbers {
		if numbers[i]%2 == 0 {
			numbers[i] = 1
		}
		sum += numbers[i]
	}

	fmt.Println("New:", numbers)
	fmt.Println("Sum:", sum)

}