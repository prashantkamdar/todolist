package main

import (
	"net/http"
		
	"github.com/gorilla/mux"
	"github.com/prashantkamdar/todolist/src/DL"
)

func main() {
	dbsetup.InitialSetup()
	router := mux.NewRouter()
	router.HandleFunc("/todolist", GetTodoList).Methods("GET")
	router.HandleFunc("/todolist", CreateTodoListItem).Methods("POST")
	router.HandleFunc("/todolist/{id}", DeleteTodoList).Methods("DELETE")
	router.HandleFunc("/todos/{todolistid}", CreateTask).Methods("POST")
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", UpdateTodo).Methods("PUT")
	http.ListenAndServe(":1337", router)
}