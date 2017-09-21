package main

import (
	"net/http"

	"github.com/peterj/goservice/movies"

	"fmt"

	"log"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	randomID := movies.GetRandomID()
	fmt.Fprintf(w, "Random ID: %s", randomID)
	// movie, err := movies.GetMovieByID(randomID)
	// if err != nil {
	// 	fmt.Fprintf(w, movie.Title)
	// } else {
	// 	log.Println("HELLO!")
	// 	fmt.Fprintf(w, err.Error())
	// }
}
func main() {
	http.HandleFunc("/api/movies", getMovies)
	fmt.Println("Listening on 5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
