package data

import (
	"time"
)

type Movie struct {
	ID        int       `json:"id"` // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`  // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   Runtime   `json:"runtime,omitzero,string"`
	Genres    []string  `json:"genres,omitzero"`
	Version   int       `json:"version"`
}
