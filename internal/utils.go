package internal

import (
	"fmt"
	"io"
	"net/http"
)

// MovieType represents the type of movie list to fetch
type MovieType string

const (
	Playing  MovieType = "playing"
	Popular  MovieType = "popular"
	TopRated MovieType = "top"
	Upcoming MovieType = "upcoming"
)

var (
	baseURL = "https://api.themoviedb.org/3/movie"
	urls    = map[MovieType]string{
		Playing:  "/now_playing",
		Popular:  "/popular",
		TopRated: "/top_rated",
		Upcoming: "/upcoming",
	}
)

// Print in a beautifull way
func PrintResult(msg string) {
	fmt.Printf("[✔] %s\n", msg)
}

// FetchMovie retrieves movie data from TMDB API based on the movie type
func FetchMovie(movieType MovieType) ([]byte, error) {
	var token = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJjMjNiM2Y3YjZlODU0YTc3MmY2YzA5M2VlMzViNzhhMCIsIm5iZiI6MTc0ODU0NTQ4NC40MjgsInN1YiI6IjY4MzhhZmNjNmJiNzJmOTg2NDA1Njk5NyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.Tx8J2pdWWTL82BnmTdOK4KMmidiO3RRLMhRIvYe2PiI"
	if token == "" {
		return nil, fmt.Errorf("tmdb token environment variable is not set")
	}

	url := baseURL + urls[movieType]
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request configuration error: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: %s", res.Status)
	}

	return io.ReadAll(res.Body)
}
