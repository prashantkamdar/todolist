package dbsetup

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type TodoList struct{
    Id int `json:"id"`
    Listname string `json:"listName"`
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