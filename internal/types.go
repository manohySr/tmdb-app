package internal

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
