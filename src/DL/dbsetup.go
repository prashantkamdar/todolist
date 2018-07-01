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

func InitialSetup(){
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error!=nil{
        log.Fatal(error)
    }
    
    statement, error := database.Prepare("CREATE TABLE IF NOT EXISTS todolist (id INTEGER PRIMARY KEY AUTOINCREMENT, listname TEXT)")
    statement.Exec()
    if error!=nil{
        log.Fatal(error)
    }

    statement, error = database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, todolistid INTEGER, taskname TEXT, FOREIGN KEY(todolistid) REFERENCES todolist(id))")
    statement.Exec()
    if error!=nil{
        log.Fatal(error)
    }

    // statement, _ = database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    // statement.Exec("firsttask")
}

func GetTodoList() ([]TodoList, error) {
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error!=nil{
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

func CreateTodoList(todolistitem TodoList) {
    database, error := sql.Open("sqlite3", "./todolist.db")
    if error!=nil{
        log.Fatal(error)
    }

    statement, _ := database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    statement.Exec(todolistitem.Listname)
}