// base command
package cmd

import (
	"fmt"
	"os"

	"github.com/manohySr/tmdb-app/internal"
	"github.com/spf13/cobra"
)

// trace prints a debug message
func trace(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}

var actionType string

var rootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "tmdb-app fetch movie from https://developer.themoviedb.org/",
	Long: `tmdb-app is a command-line tool that lets you fetch and interact with
movie data from The Movie Database (TMDB).`,
	Run: func(cmd *cobra.Command, args []string) {
		switch actionType {
		case "playing":
			trace("entering in playing")
			response, err := internal.FetchMovie(internal.Playing)
			if err != nil {
				internal.PrintResult(fmt.Sprintf("Error: %v", err))
				return
			}

			internal.PrintResult(fmt.Sprintf("Found %d movies now playing:", len(response.Results)))
			for _, movie := range response.Results {
				internal.PrintMovie(movie)
			}
		case "popular":
			response, err := internal.FetchMovie(internal.Popular)
			if err != nil {
				internal.PrintResult(fmt.Sprintf("Error: %v", err))
				return
			}
			internal.PrintResult(fmt.Sprintf("Found %d popular movies:", len(response.Results)))
			for _, movie := range response.Results {
				internal.PrintMovie(movie)
			}
		case "top":
			response, err := internal.FetchMovie(internal.TopRated)
			if err != nil {
				internal.PrintResult(fmt.Sprintf("Error: %v", err))
				return
			}
			internal.PrintResult(fmt.Sprintf("Found %d top rated movies:", len(response.Results)))
			for _, movie := range response.Results {
				internal.PrintMovie(movie)
			}
		case "upcoming":
			response, err := internal.FetchMovie(internal.Upcoming)
			if err != nil {
				internal.PrintResult(fmt.Sprintf("Error: %v", err))
				return
			}
			internal.PrintResult(fmt.Sprintf("Found %d upcoming movies:", len(response.Results)))
			for _, movie := range response.Results {
				internal.PrintMovie(movie)
			}
		default:
			internal.PrintResult("Error: You must provide a valid --type (playing, popular, top, upcoming)")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&actionType, "type", "", "Specify the action type (playing, popular, top, upcoming)")
}
