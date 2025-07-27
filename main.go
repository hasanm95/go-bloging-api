package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
