# 🎬 TMDB CLI Tool

A beautiful command-line interface tool to fetch and display movie information from The Movie Database (TMDB) API. This project was inspired by [roadmap.sh TMDB CLI project](https://roadmap.sh/projects/tmdb-cli).

## ✨ Features

- Fetch and display movies by different categories:
  - Now Playing Movies
  - Popular Movies
  - Top Rated Movies
  - Upcoming Movies
- Beautiful terminal output with color-coded information
- Detailed movie information including:
  - Title
  - Release Date
  - Rating
  - Vote Count
  - Overview

## 🚀 Installation

### From GitHub

```bash
# Install the latest version
go install github.com/manohySr/tmdb-app@latest
```

### From Source

```bash
# Clone the repository
git clone https://github.com/manohySr/tmdb-app.git

# Navigate to the project directory
cd tmdb-app

# Build the project
go build -o tmdb-app

# Move to a directory in your PATH (optional)
sudo mv tmdb-app /usr/local/bin/
```

## 💻 Usage

```bash
# Show now playing movies
tmdb-app --type "playing"

# Show popular movies
tmdb-app --type "popular"

# Show top rated movies
tmdb-app --type "top"

# Show upcoming movies
tmdb-app --type "upcoming"
```

## 📁 Project Structure

```
https://gitingest.com/manohySr/tmdb-app.git
```

## 📋 Requirements

- Go 1.23 or higher
- TMDB API - Powered by [The Movie Database](https://www.themoviedb.org/)

## 🤝 Contributing

Feel free to submit issues and enhancement requests! We welcome all contributions.

## 📝 License

This project is open source and available under the MIT License.

---
Made with ❤️ by [manohySr](https://github.com/manohySr)