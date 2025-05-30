package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchMovie retrieves movie data from TMDB API based on the movie type
func FetchMovie(movieType MovieType) (*MovieResponse, error) {
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response MovieResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}
