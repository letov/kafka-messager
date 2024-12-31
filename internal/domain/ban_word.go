package domain

import "time"

type BanWord struct {
	Id        int64     `json:"id"`
	Word      int64     `json:"word"`
	CreatedAt time.Time `json:"created_at"`
}
