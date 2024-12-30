package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
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
)

func main() {
	http.HandleFunc("/apply", applyHandler)
	http.HandleFunc("/admitted", admittedHandler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func applyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	totalScore := student.MathScore + student.InformaticsScore + student.EnglishScore
	if totalScore >= 14 {
		mu.Lock()
		admittedStudents = append(admittedStudents, student)
		mu.Unlock()
		fmt.Fprintf(w, "Student %s admitted\n", student.FullName)
	} else {
		fmt.Fprintf(w, "Student %s not admitted\n", student.FullName)
	}
}

func admittedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admittedStudents)
}
