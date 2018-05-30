package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("client/build")))
    log.Fatal(http.ListenAndServe(":8000", nil))
}
