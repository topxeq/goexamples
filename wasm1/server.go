package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8835", http.FileServer(http.Dir("."))))
}
