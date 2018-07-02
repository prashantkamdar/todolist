package main

import (
	"encoding/json"
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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
	todoIdToDelete,_ :=  strconv.Atoi(params["id"])
	dbsetup.DeleteTodo(todoIdToDelete)
	json.NewEncoder(w).Encode("success")
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	var todo dbsetup.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	dbsetup.UpdateTodo(todo)
	json.NewEncoder(w).Encode("success")
}