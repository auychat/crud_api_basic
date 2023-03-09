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
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

// สร้าง slices
var movies []Movie

// สร้าง getMovies Function สำหรับ GET ALL Routes
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// สร้าง getMovie Function สำหรับ GET BY ID Routes
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// สร้าง createMovie Function สำหรับ CREATE Routes
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// สร้าง updateMovie Function สำหรับ UPDATE Routes
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//loop over the movies, range
	//delete the movie with the i.d that you've sent
	//add a new movie - the movie that we send in the body of postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

// สร้าง deleteMovie Function สำหรับ DELETE Routes
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	// ใส่ Initial value ให้กับ Movies slices
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// สร้าง 5 Routes ตามแผนผัง GET ALL, GET BY ID, CREATE, UPDATE, DELETE
	//GET ALL
	r.HandleFunc("/movies", getMovies).Methods("GET")

	//GET BY ID
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	//CREATE
	r.HandleFunc("/movies", createMovie).Methods("POST")

	//UPDATE
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	//DELETE
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
