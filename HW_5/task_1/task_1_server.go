// Цель: Создать HTTP-сервер на языке Go, который будет обрабатывать заявки студентов на поступление в университет.
// Сервер должен принимать данные о студентах, проверять их баллы и выводить список поступивших студентов.

// Задачи:
// Создание структуры данных:
// Определите структуру Student, которая будет содержать следующие поля:
// FullName (строка) — полное имя студента.
// MathScore (целое число) — балл по математике.
// InformaticsScore (целое число) — балл по информатике.
// EnglishScore (целое число) — балл по английскому языку.
// Создание HTTP-сервера:
// Реализуйте HTTP-сервер, который будет слушать на порту 8080.
// Обработчик для поступления:
// Создайте обработчик для POST-запросов на маршрут /apply, который будет принимать JSON с данными студента.
// В обработчике проверьте, если сумма баллов по трем предметам (математика, информатика, английский)
// больше или равна 14, то добавьте студента в список поступивших. В противном случае, верните сообщение о том,
// что студент не поступил.
// Создание студентов:
// Создайте трех студентов (клиентов) с заранее определенными баллами:
// Два студента должны иметь общую сумму баллов >= 14.
// Один студент должен иметь общую сумму баллов < 14.
// Обработчик для вывода поступивших студентов:
// Создайте новый маршрут /admitted, который будет возвращать список всех студентов, которые поступили.
// Список должен быть представлен в формате JSON.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type Student struct {
	FullName         string `json:"fullName"`
	MathScore        int    `json:"mathScore"`
	InformaticsScore int    `json:"informaticsScore"`
	EnglishScore     int    `json:"englishScore"`
}

var (
	admittedStudents []Student
	mu               sync.Mutex
	log              = logrus.New()
)

func main() {
	http.HandleFunc("/apply", applyHandler)
	http.HandleFunc("/admitted", admittedHandler)

	log.Info("Server is starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func applyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postStudent(w, r)
	default:
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Warn("Invalid method")
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func postStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	totalScore := student.MathScore + student.InformaticsScore + student.EnglishScore
	if totalScore >= 14 {
		mu.Lock()
		admittedStudents = append(admittedStudents, student)
		mu.Unlock()
		log.WithFields(logrus.Fields{
			"student": student.FullName,
		}).Info("Student admitted")
		fmt.Fprintf(w, "Student %s admitted\n", student.FullName)
	} else {
		log.WithFields(logrus.Fields{
			"student": student.FullName,
		}).Info("Student not admitted")
		fmt.Fprintf(w, "Student %s not admitted\n", student.FullName)
	}
}

func admittedHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAdmittedStudents(w)
	default:
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Warn("Invalid method")
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func getAdmittedStudents(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(admittedStudents)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to encode admitted students")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("Returned list of admitted students")
}
