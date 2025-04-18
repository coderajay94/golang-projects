package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/coderajay94/bookservice-go/model"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := model.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal and send the response as JSON
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode books", http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := model.DeleteBook(id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Book deleted successfully")
	w.WriteHeader(http.StatusOK)

}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := model.GetBookByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := book.CreateBook(); err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook model.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := updatedBook.UpdateBook(id); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return
// 	}

// 	if err := model.DeleteBook(id); err != nil {
// 		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
