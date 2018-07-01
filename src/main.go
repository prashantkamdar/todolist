package main

import (
    "net/http"
    "log"	
    "github.com/gorilla/mux"
)

// our main function
func main() {
    router := mux.NewRouter()
    log.Fatal(http.ListenAndServe(":1337", router))
}