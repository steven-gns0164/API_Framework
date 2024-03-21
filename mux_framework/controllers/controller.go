package controllers

import (
	"encoding/json"
	"log"
	m "mux_framework/models"
	"net/http"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM students"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to execute query"})
		return
	}
	defer rows.Close()

	var students []m.Students
	for rows.Next() {
		var student m.Students
		if err := rows.Scan(&student.NIM, &student.Nama); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve data"})
			return
		}
		students = append(students, student)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  200,
		"message": "Success",
		"data":    students,
	})

	SendSuccessResponse(w, http.StatusOK, "Success")
}

func InsertNewStudent(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse request body
	NIM := r.URL.Query().Get("nim")
	Nama := r.URL.Query().Get("nama")

	// Insert the new student into the database
	query := "INSERT INTO students (NIM, Nama) VALUES (?, ?)"
	_, err := db.Exec(query, NIM, Nama)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to insert new student")
		return
	}
	SendSuccessResponse(w, http.StatusCreated, "Student inserted successfully")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse request body
	NIM := r.URL.Query().Get("nim")
	Nama := r.URL.Query().Get("nama")

	// Update the student information in the database
	query := "UPDATE students SET Nama=? WHERE NIM=?"
	_, err := db.Exec(query, Nama, NIM)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to update student information")
		return
	}
	SendSuccessResponse(w, http.StatusOK, "Student information updated successfully")
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse student ID from the URL query parameters
	NIM := r.URL.Query().Get("nim")
	if NIM == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Student NIM is required")
		return
	}

	// Execute the SQL query to delete the student record
	query := "DELETE FROM students WHERE NIM = ?"
	_, err := db.Exec(query, NIM)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	// Send success response if deletion is successful
	SendSuccessResponse(w, http.StatusOK, "Student deleted successfully")
}
