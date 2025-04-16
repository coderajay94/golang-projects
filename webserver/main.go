package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error Starting server", err)
	}
}
