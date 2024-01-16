package types

import (
	"time"
)

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	Isbn        int64     `json:"isbn"`
	Available   int       `json:"available"`
	Added_At    time.Time `json:"added_at"`
}

type BookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Available   int    `json:"available"`
	Isbn        int64  `json:"isbn"`
}

type BookResponse struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	Isbn        int64     `json:"isbn"`
	Available   int       `json:"available"`
	Added_At    time.Time `json:"added_at"`
}
