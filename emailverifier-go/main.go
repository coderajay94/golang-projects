package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type EmailResponse struct {
	IsValid bool `json:"isValid"`
}

func main() {
	fmt.Println("starting the server....")

	r := mux.NewRouter()
	r.HandleFunc("/validate/{emailId}", EmailValidator).Methods("GET")
	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func EmailValidator(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	emailId := params["emailId"]
	w.Header().Set("Content-Type", "application/json")

	if emailId == "" {
		json.NewEncoder(w).Encode(EmailResponse{IsValid: false})
	}

	if resp, err := net.LookupTXT(emailId); len(resp) == 0 && err != nil {
		json.NewEncoder(w).Encode(EmailResponse{IsValid: false})
	}
	json.NewEncoder(w).Encode(EmailResponse{IsValid: true})
	fmt.Println("Email is valid")
}
