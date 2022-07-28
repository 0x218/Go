package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var sliceMovies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sliceMovies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range sliceMovies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tmpMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&tmpMovie)
	tmpMovie.ID = strconv.Itoa(rand.Intn(1000000))
	sliceMovies = append(sliceMovies, tmpMovie)

	json.NewEncoder(w).Encode(tmpMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//delete given movie and add the movie with updated info.
	params := mux.Vars(r)
	for index, item := range sliceMovies {
		if item.ID == params["id"] {
			sliceMovies = append(sliceMovies[:index], sliceMovies[index+1:]...) //copy from index position to end, thus remove the index position data

			var tmpMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&tmpMovie)
			tmpMovie.ID = params["id"]
			sliceMovies = append(sliceMovies, tmpMovie)
			json.NewEncoder(w).Encode(tmpMovie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range sliceMovies {
		if item.ID == params["id"] {
			sliceMovies = append(sliceMovies[:index], sliceMovies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(sliceMovies)
}

func main() {

	sliceMovies = append(sliceMovies, Movie{ID: "1", ISBN: "23435", Title: "Movie# 1", Director: &Director{FirstName: "Pandit", LastName: "Bakkar"}})
	sliceMovies = append(sliceMovies, Movie{ID: "2", ISBN: "34325", Title: "Movie #2", Director: &Director{FirstName: "Kishore", LastName: "Chakra"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
