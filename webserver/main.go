package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error Starting server", err)
	}
	fmt.Println("App is running on port 8080")

}

func formHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Form Handler")
	if err := request.ParseForm(); err != nil {
		fmt.Println("Error parsing form", err)
		return
	}

	fmt.Println("Form data received")
	fmt.Println("Name:", request.FormValue("name"))
	fmt.Println("Email:", request.FormValue("email"))
	fmt.Println("Phone:", request.FormValue("phone"))
	fmt.Println("Gender:", request.FormValue("gender"))
	fmt.Fprintf(response, "Form data Processed Successfully for "+request.FormValue("name"))

}
func helloHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello Handler")
	if request.Method == http.MethodPost {
		fmt.Println("post method is not allowed")
		return
	}

	fmt.Fprintf(response, "Hello")
}
