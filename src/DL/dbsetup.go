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
    
    statement, error := database.Prepare(`CREATE TABLE IF NOT EXISTS todolist
        (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            listname TEXT,
            updatedate DATETIME DEFAULT CURRENT_TIMESTAMP
        )`)
    statement.Exec()
    if error != nil{
        log.Fatal(error)
    }

    statement, error = database.Prepare(`CREATE TABLE IF NOT EXISTS todos
        (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            todolistid INTEGER,
            taskname TEXT,
            updatedate DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(todolistid) REFERENCES todolist(id)
        )`)
    statement.Exec()
    if error != nil{
        log.Fatal(error)
    }
    
    //dummy data
    // statement, _ = database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    // statement.Exec("firstlist")
    // statement, _ = database.Prepare("INSERT INTO todolist (listname) VALUES (?)")
    // statement.Exec("secondlist")
    // statement, _ = database.Prepare("INSERT INTO todos (todolistid,taskname) VALUES (?,?)")
    // statement.Exec(1, "firstlistfirsttask")
    // statement, _ = database.Prepare("INSERT INTO todos (todolistid,taskname) VALUES (?,?)")
    // statement.Exec(1, "firstlistsecondtask")
    // statement, _ = database.Prepare("INSERT INTO todos (todolistid,taskname) VALUES (?,?)")
    // statement.Exec(2, "secondlistfirsttask")
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