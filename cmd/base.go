// base command
package cmd

import (
	"fmt"
	"os"

	"github.com/manohySr/tmdb-app/internal"
	"github.com/spf13/cobra"
)

var actionType string

var rootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "tmdb-app fetch movie from https://developer.themoviedb.org/",
	Long: `tmdb-app is a command-line tool that lets you fetch and interact with
movie data from The Movie Database (TMDB).`,
	Run: func(cmd *cobra.Command, args []string) {
		switch actionType {
		case "playing":
			data, err := internal.FetchMovie(internal.Playing)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Success! Got data: %s\n", string(data))
		case "popular":
			internal.PrintResult("popular...")
		case "top":
			internal.PrintResult("top...")
		case "upcoming":
			internal.PrintResult("upcoming...")
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
