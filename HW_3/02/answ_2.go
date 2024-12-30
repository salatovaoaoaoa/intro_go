// Почему в одной версии go 1.20 выводится 40 40 40 40, а в другой 1.23 выводится 10 20 30 40

package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &value)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}

// Если коротко: поменяли поведение компилятора и оптимизации в версии 1.22 (https://tip.golang.org/doc/go1.22):

// Previously, the variables declared by a “for” loop were created once and updated by each iteration.
// In Go 1.22, each iteration of the loop creates new variables, to avoid accidental sharing bugs.
// The transition support tooling described in the proposal continues to work in the same way it did in Go 1.21.

// Подробнее
// Go 1.20: Переменная value в цикле for создаётся один раз и обновляется на каждой итерации.
// При добавлении &value в срез numbers каждый элемент ссылается на одну и ту же переменную value,
// которая после последней итерации содержит значение 40. Поэтому при выводе отображается 40 40 40 40.

// Go 1.22: Теперь на каждой итерации цикла создаётся новая переменная value.
// При добавлении &value в срез numbers каждый элемент ссылается на уникальную переменную,
// содержащую соответствующее значение из массива [10, 20, 30, 40]. Поэтому при выводе отображается 10 20 30 40.
