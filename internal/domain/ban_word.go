package domain

import "time"

type BanWord struct {
	Id        int64     `json:"id"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"created_at"`
}
