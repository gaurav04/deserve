package main

import (
	"log"
	"net/http"
	handlers "talentica/handlers"
)

func handleRequests() {
	http.HandleFunc("/", handlers.JsonHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
