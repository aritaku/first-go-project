package main

import (
    "log"
		"net/http"

		"github.com/aritaku/first-go-project/app"
		"github.com/aritaku/first-go-project/config"
)

// main function to boot up everything
func main() {
		router := NewRouter()

    log.Fatal(http.ListenAndServe(":8000", router))
}
