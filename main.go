package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", apiHandler)

	log.Println("Server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling api handler")
}
