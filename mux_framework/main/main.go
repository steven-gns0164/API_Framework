package main

import (
	"fmt"
	"log"
	"net/http"

	"mux_framework/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/controller", controllers.GetAllStudents).Methods("GET")
	r.HandleFunc("/controller", controllers.InsertNewStudent).Methods("POST")
	r.HandleFunc("/controller/", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/controller", controllers.DeleteStudent).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
