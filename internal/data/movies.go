package data

import (
	"time"
)

type Movie struct {
	ID        int       // Unique integer ID for the movie
	CreatedAt time.Time // Timestamp for when the movie is added to our database
	Title     string
	Year      int
	Runtime   int
	Genres    []string
	Version   int
}
