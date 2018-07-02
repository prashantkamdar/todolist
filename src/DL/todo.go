package dbsetup

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Todo struct{
    Id int `json:"id"`
    Todolistid int `json:"todoListId"`
    Taskname string `json:"taskName"`
}

func CreateTask(todo Todo){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }

    statement, error := database.Prepare("INSERT INTO todos (todolistid, taskname) VALUES (?, ?)")
    statement.Exec(todo.Todolistid, todo.Taskname)
    if error != nil{
        log.Fatal(error)
    }
}

func GetTodos() ([]Todo, error) {
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
        return nil, error
    }
    var todos []Todo

    rows, error := database.Query("SELECT id, todolistid, taskname FROM todos")
    var id int
    var todolistid int
    var taskname string
    var todoitem Todo

    for rows.Next() {
        rows.Scan(&id, &todolistid, &taskname)
        todoitem.Id = id
        todoitem.Todolistid = todolistid
        todoitem.Taskname = taskname

        todos = append(todos, todoitem)
    }

    return todos, nil
}

func DeleteTodo(todoIdToDelete int){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }

    statement, error := database.Prepare("DELETE FROM todos WHERE id = ?")
    statement.Exec(todoIdToDelete)
    if error != nil{
        log.Fatal(error)
    }
}

func UpdateTodo(todo Todo){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }

    statement, error := database.Prepare("UPDATE todos SET todolistid = ?, taskname = ? WHERE id = ?")
    statement.Exec(todo.Todolistid, todo.Taskname, todo.Id)
    if error != nil{
        log.Fatal(error)
    }
}