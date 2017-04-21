package main

import (
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PeterJausovec/goservice/movies"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	movie, err := movies.GetMovieByID(randomID())
	if err != nil {
		io.WriteString(w, movie.Title)
	}
}
func main() {
	http.HandleFunc("/api/movies", getMovies)
	http.ListenAndServe(":5000", nil)
}

func randomID() string {
	rand.Seed(time.Now().UnixNano())
	res := "tt1"
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(9))
	}
	return res
}
