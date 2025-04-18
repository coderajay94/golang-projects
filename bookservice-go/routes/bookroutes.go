package routes

import (
	"log"

	"github.com/coderajay94/bookservice-go/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(r *mux.Router) {
	log.Println("RegisterBookRoutes called")
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
