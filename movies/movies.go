package movies

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	omdbURL = "http://www.omdbapi.com/?i="
)

// MovieData contains the info returned from calls to OMDBApi.
type MovieData struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Season     string `json:"Season"`
	Episode    string `json:"Episode"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	SeriesID   string `json:"seriesID"`
	Type       string `json:"Type"`
	Response   string `json:"Response"`
}

// GetMovieByID makes a request to OMDB and returns a movie using the imdb ID.
func GetMovieByID(imdbID string) (MovieData, error) {
	url := omdbURL + imdbID
	resp, err := http.Get(url)

	if err != nil {
		return MovieData{}, err
	}
	defer resp.Body.Close()
	var m MovieData
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return MovieData{}, err
	}

	return m, nil
}

func GetRandomID() string {
	rand.Seed(time.Now().UnixNano())
	res := "tt1"
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(9))
	}
	return res
}
