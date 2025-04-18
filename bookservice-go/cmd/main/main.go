package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coderajay94/bookservice-go/model"
	"github.com/coderajay94/bookservice-go/routes"
	"github.com/gorilla/mux"
)

func main() {

	model.Init()
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started at 8080 successfully!")
}
