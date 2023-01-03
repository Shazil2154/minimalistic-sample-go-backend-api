package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"_"`
	UpdatedAt   time.Time      `json:"_"`
	MovieGenre  map[int]string `json:"movie_genre"`
}

type Genre struct {
	ID        int       `json:"_"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

type MovieGenre struct {
	ID        int       `json:"_"`
	MovieID   int       `json:"_"`
	GenreID   int       `json:"_"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

type User struct {
	ID       int
	Email    string
	Password string
}
