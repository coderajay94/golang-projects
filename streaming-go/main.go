package main

import (
	"fmt"
	"log"
	"net/http"
	"streaming-go/handlers"
)

func main() {
	fmt.Println("Streaming starts from here...")

	http.HandleFunc("/events", handlers.StreamTime)

	directory := http.Dir("./static")
	fs := http.FileServer(directory)

	http.Handle("/", fs)

	fmt.Println("Server started on port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
