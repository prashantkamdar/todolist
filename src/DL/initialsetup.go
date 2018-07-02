package dbsetup

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

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