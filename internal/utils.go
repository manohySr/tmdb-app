package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

// Movie represents a movie from TMDB API
type Movie struct {
	OriginalTitle string  `json:"original_title"`
	Title         string  `json:"title"`
	ReleaseDate   string  `json:"release_date"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     float64 `json:"vote_count"`
	Overview      string  `json:"overview"`
}

// MovieResponse represents the TMDB API response
type MovieResponse struct {
	Results []Movie `json:"results"`
}

// Print in a beautifull way
func PrintResult(msg string) {
	fmt.Printf("\033[1;32m[✔]\033[0m %s\n", msg)
}

// PrintMovie prints movie information in a beautiful format
func PrintMovie(movie Movie) {
	// Format the overview to wrap at 80 characters
	overview := wrapText(movie.Overview, 80)

	fmt.Printf("\n\033[1;36m%s\033[0m\n", strings.Repeat("─", 90))
	fmt.Printf("\033[1;33m🎬 %s\033[0m\n", movie.Title)
	fmt.Printf("\033[1;32m📅 Release Date:\033[0m %s\n", movie.ReleaseDate)
	fmt.Printf("\033[1;35m⭐ Rating:\033[0m %.1f/10 (\033[1;34m%v votes\033[0m)\n", movie.VoteAverage, movie.VoteCount)
	fmt.Printf("\033[1;31m📝 Overview:\033[0m\n%s\n", overview)
	fmt.Printf("\033[1;36m%s\033[0m\n", strings.Repeat("─", 90))
}

// wrapText wraps text at the specified width
func wrapText(text string, width int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var lines []string
	currentLine := "\t"
	currentLength := 0

	for _, word := range words {
		if currentLength+len(word)+1 > width {
			lines = append(lines, currentLine)
			currentLine = "\t" + word
			currentLength = len(word)
		} else {
			if currentLength > 0 {
				currentLine += " "
				currentLength++
			}
			currentLine += word
			currentLength += len(word)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

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
