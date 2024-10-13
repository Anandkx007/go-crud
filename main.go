package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.Id == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.IntN(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.Id == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.IntN(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "289782", Title: "Movie1", Director: &Director{FistName: "Anand", LastName: "kumar"}})
	movies = append(movies, Movie{Id: "2", Isbn: "289783", Title: "Movie2", Director: &Director{FistName: "Aman", LastName: "kumar"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8082 \n")

	log.Fatal(http.ListenAndServe(":8082", r))

}
