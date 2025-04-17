package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

type Movie struct {
	Title    string    `json:"title"`
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(Response{Status: "Failure", Description: "Item not found."})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(Response{Status: "Success", Description: "Item deleted Successfully."})
			return
		}
	}
	json.NewEncoder(w).Encode(Response{Status: "Failure", Description: "Item not found."})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		json.NewEncoder(w).Encode(Response{Status: "Failure", Description: "Invalid request payload."})
		return
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(Response{Status: "Success", Description: "Movie Created Successfully."})
}

func gethealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Status: "Up & Running..."})
}

func main() {

	movies = append(movies, Movie{Title: "Matrix", ID: "1", ISBN: "438383", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{Title: "Titanic", ID: "2", ISBN: "438663", Director: &Director{FirstName: "Ajay", LastName: "Kumar"}})
	movies = append(movies, Movie{Title: "Titans", ID: "5", ISBN: "438783", Director: &Director{FirstName: "Lwa", LastName: "sue"}})
	movies = append(movies, Movie{Title: "Titanic", ID: "2", ISBN: "438683", Director: &Director{FirstName: "Ajay", LastName: "Kumar"}})
	movies = append(movies, Movie{Title: "Honda", ID: "4", ISBN: "446683", Director: &Director{FirstName: "lou", LastName: "Nar"}})

	fmt.Println("Hello, World!")
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/health", gethealth).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movie/{id}", updateMovie).Methods("PUT")
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
