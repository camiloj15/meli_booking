package model

import "time"

type Book struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int64     `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	Ranking   float64   `json:"ranking"`
}
