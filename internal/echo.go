package internal

import (
	"fmt"
	"strings"
)

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

// wraps text at the specified width
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
