// Задача 1:
// Описание задачи:
// Создайте программу для управления библиотекой.
// Каждый экземпляр книги должен иметь название, автора, год издания и статус
// (доступна или на руках у читателя).
// Добавьте возможность добавления новой книги, поиска книги по названию, выдачи книги читателю и возврата книги.

// Требования:

// Используйте структуры для представления книги.
// Напишите методы для добавления книги, выдачи книги читателю и возврата книги.
// Используйте функции для поиска книги по названию.
// Используйте циклы для вывода списка всех книг.
// Подсказка:

// Метод Issue меняет статус книги на "на руках у читателя".
// Метод Return меняет статус книги на "доступна".

package main

import (
	"fmt"
	"strings"
)

// структура для книги
type Book struct {
	Title  string
	Author string
	Year   int
	Status string
}

// структура библиотека для структуры книга
type Library struct {
	Books []Book
}

// метод добавления новой книги в библиотеку
func (l *Library) AddingBook(title, author string, year int) {
	book := Book{
		Title:  title,
		Author: author,
		Year:   year,
		Status: "доступна",
	}
	l.Books = append(l.Books, book)
	fmt.Println("Добавлена книга:", book.Title)
}

// метод ищет книгу по названию и возвращает указатель на неё
func (l *Library) FindBook(title string) *Book {
	for i := range l.Books {
		if strings.EqualFold(l.Books[i].Title, title) {
			return &l.Books[i]
		}
	}
	return nil
}

// метод для выдачи книги (меняем статус)
func (b *Book) Issue() {
	if b.Status == "доступна" {
		b.Status = "на руках у читателя"
		fmt.Println("Книга выдана:", b.Title)
	} else {
		fmt.Println("Книга не доступна (находится 'на руках'):", b.Title)
	}
}

// метод для возврата книги
func (b *Book) Return() {
	if b.Status == "на руках у читателя" {
		b.Status = "доступна"
		fmt.Println("Книга возвращена:", b.Title)
	} else {
		fmt.Println("Книга находится в библиотеке:", b.Title)
	}
}

// метод выводит список всех книг в библиотеке
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("В библиотеке нет книг.")
		return
	}

	fmt.Println("Список книг в библиотеке:")
	for _, book := range l.Books {
		fmt.Printf("Название: %s, Автор: %s, Год: %d, Статус: %s\n",
			book.Title, book.Author, book.Year, book.Status)
	}
}

func main() {
	library := Library{}

	// Внесение книг
	library.AddingBook("Игрок", "Ф.М. Достоевский", 1866)
	library.AddingBook("Анна Каренина", "Л. Н. Толстой", 1873)
	library.AddingBook("Преступление и наказание", "Ф.М. Достоевский", 1866)
	library.AddingBook("Идиот", "Ф.М. Достоевский", 1868)
	library.AddingBook("Морфий", "М. Булгаков", 1926)

	library.ListBooks()

	book := library.FindBook("Игрок")
	if book != nil {
		book.Issue()
	} else {
		fmt.Println("Книга не найдена")
	}

	if book != nil {
		book.Issue()
	}

	if book != nil {
		book.Return()
	}

	library.ListBooks()

	// если такой книги нет в библиотеке
	missingBook := library.FindBook("Белые ночи")
	if missingBook != nil {
		missingBook.Issue()
	} else {
		fmt.Println("Книга не найдена")
	}
}
