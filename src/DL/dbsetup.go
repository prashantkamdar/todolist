package dbsetup

import (
	//"errors"
	//"fmt"    
    //"strconv"
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type TodoList struct{
    Id int `json:"id"`
    Listname string `json:"listName"`
}

type Todo struct{
    Id int `json:"id"`
    Todolistid int `json:"todoListId"`
    Taskname string `json:"taskName"`
}

func InitialSetup(){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }
    
    statement, error := database.Prepare("CREATE TABLE IF NOT EXISTS todolist (id INTEGER PRIMARY KEY AUTOINCREMENT, listname TEXT)")
    statement.Exec()
    if error != nil{
        log.Fatal(error)
    }

    statement, error = database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, todolistid INTEGER, taskname TEXT, FOREIGN KEY(todolistid) REFERENCES todolist(id))")
    statement.Exec()
    if error != nil{
        log.Fatal(error)
    }

    // statement, _ = database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    // statement.Exec("firsttask")
}

func GetTodoList() ([]TodoList, error) {
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
        return nil, error
    }
    var todolist []TodoList

    rows, error := database.Query("SELECT id, listname FROM todolist")
    var id int
    var listname string
    var todolistitem TodoList

    for rows.Next() {
        rows.Scan(&id, &listname)
        todolistitem.Id = id
        todolistitem.Listname = listname

        todolist = append(todolist, todolistitem)
    }

    return todolist, nil
}

func CreateTodoListItem(todolistitem TodoList) {
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }

    statement, error := database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    statement.Exec(todolistitem.Listname)
    if error != nil{
        log.Fatal(error)
    }
}

func DeleteTodoListItem(taskIdToDelete int){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error != nil{
        log.Fatal(error)
    }

    statement, error := database.Prepare("DELETE FROM todolist WHERE id = ?")
    statement.Exec(taskIdToDelete)
    if error != nil{
        log.Fatal(error)
    }
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

    rows, error := database.Query("SELECT todolistid, taskname FROM todos")
    var todolistid int
    var taskname string
    var todoitem Todo

    for rows.Next() {
        rows.Scan(&todolistid, &taskname)
        todoitem.Todolistid = todolistid
        todoitem.Taskname = taskname

        todos = append(todos, todoitem)
    }

    return todos, nil
}