// Задача 2:

// Описание задачи:

// Создайте программу, которая поможет пользователям учитывать свои ежемесячные расходы.
// Используйте карту, где ключом будет название категории расходов (например, "Продукты", "Транспорт", "Развлечения"),
//  а значением - сумма расходов по этой категории.

// Требования:
// Используя функции
// Пользователь должен иметь возможность добавлять новые категории и записывать расходы по каждой из них.
// Также добавьте функцию для подсчета общей суммы расходов и вывода ее на экран

package main

import (
	"fmt"
)

// map в структуре для учета расходов
type CostCount struct {
	categories map[string]float64
}

// метод создает новый экземпляр CostCount
func NewCostCount() *CostCount {
	return &CostCount{
		categories: make(map[string]float64),
	}
}

// метод для добавления новой категории расходов
func (e *CostCount) AddCategory(category string) {
	if _, exists := e.categories[category]; exists {
		fmt.Println("Категория уже существует:", category)
	} else {
		e.categories[category] = 0
		fmt.Println("Добавлена новая категория:", category)
	}
}

// метод для добавления расхода в указанную категорию
func (e *CostCount) AddExpense(category string, amount float64) {
	if _, exists := e.categories[category]; exists {
		e.categories[category] += amount
		fmt.Printf("Добавлено %.1f в категорию %s\n", amount, category)
	} else {
		fmt.Println("Категория не найдена:", category)
	}
}

// общая сумма
func (e *CostCount) GetTotal() float64 {
	total := 0.0
	for _, amount := range e.categories {
		total += amount
	}
	return total
}

// отчет расходов по категориям
func (e *CostCount) PrintSummary() {
	fmt.Println("Отчет расходов по категориям:")
	for category, amount := range e.categories {
		fmt.Printf("%s: %.1f\n", category, amount)
	}
	total := e.GetTotal()
	fmt.Printf("Всего потрачено: %.1f\n", total)
}

func main() {
	expenseTracker := NewCostCount()

	// Добавление категорий
	expenseTracker.AddCategory("Продукты")
	expenseTracker.AddCategory("Транспорт")
	expenseTracker.AddCategory("Развлечения")
	expenseTracker.AddCategory("Косметика")
	expenseTracker.AddCategory("WB")

	// Добавление расходов
	expenseTracker.AddExpense("Продукты", 1499.9)
	expenseTracker.AddExpense("Транспорт", 1305.0)
	expenseTracker.AddExpense("Развлечения", 1800.0)
	expenseTracker.AddExpense("Косметика", 4800.0)
	expenseTracker.AddExpense("WB", 2787.4)

	// Вывод отчета
	expenseTracker.PrintSummary()

	// Если категория не существует
	expenseTracker.AddExpense("Рестики", 3000.00)

	// Добавление новой категории и расходов
	expenseTracker.AddCategory("Рестики")
	expenseTracker.AddExpense("Рестики", 3000.00)

	// Вывод отчета после обновления
	expenseTracker.PrintSummary()
}
