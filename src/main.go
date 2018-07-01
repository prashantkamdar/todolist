package main

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"
		
	"github.com/gorilla/mux"
	"github.com/prashantkamdar/todolist/src/DL"
)

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	todolist, err := dbsetup.GetTodoList()
	if err != nil{
		json.NewEncoder(w).Encode("error")		
    } else{
		json.NewEncoder(w).Encode(todolist)
	}
}

func CreateTodoListItem(w http.ResponseWriter, r *http.Request) {
    var todolistitem dbsetup.TodoList
	_ = json.NewDecoder(r.Body).Decode(&todolistitem)
	dbsetup.CreateTodoListItem(todolistitem)
	json.NewEncoder(w).Encode("success")
}

func DeleteTodoList(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
	taskIdToDelete,_ :=  strconv.Atoi(params["id"])
	dbsetup.DeleteTodoListItem(taskIdToDelete)
	json.NewEncoder(w).Encode("success")
}

func CreateTask(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
	todolistid,_ :=  strconv.Atoi(params["todolistid"])
	var todo dbsetup.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.Todolistid = todolistid

	dbsetup.CreateTask(todo)
	json.NewEncoder(w).Encode("success")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := dbsetup.GetTodos()
	if err != nil{
		json.NewEncoder(w).Encode("error")		
    } else{
		json.NewEncoder(w).Encode(todos)
	}
}

func main() {
	dbsetup.InitialSetup()
	router := mux.NewRouter()
	router.HandleFunc("/todolist", GetTodoList).Methods("GET")
	router.HandleFunc("/createlist", CreateTodoListItem).Methods("POST")
	router.HandleFunc("/todolist/{id}", DeleteTodoList).Methods("DELETE")
	router.HandleFunc("/createtodo/{todolistid}", CreateTask).Methods("POST")
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	http.ListenAndServe(":1337", router)
}