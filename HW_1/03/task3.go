// Задача:
// Дан слайс чисел, необходимо найти минимальное и максимальное значение, которое делится на 2 без остатка.
// numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}

package main

import (
	"fmt"
	"math"
)

func main() {

	numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}

	min := math.MaxInt
	max := math.MinInt

	var numDevtwo []int

	for i := range numbers {
		if numbers[i]%2 == 0 {
			numDevtwo = append(numDevtwo, numbers[i])
			if numbers[i] < min {
				min = numbers[i]
			}
			if numbers[i] > max {
				max = numbers[i]
			}
		}
	}
	//дополнительно проверю, вдруг у нас все числа в массиве нечетные или вообще там нет чисел:

	if len(numDevtwo) > 0 {
		fmt.Println("Минимальное делится на 2:", min)
		fmt.Println("Макс делится на 2:", max)
	} else {
		fmt.Println("В исходном массиве нет чисел, делящихся на два")
	}
}
