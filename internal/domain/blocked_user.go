package domain

import "time"

type BlockedUser struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	BlockUserId int64     `json:"recipient_id"`
	CreatedAt   time.Time `json:"created_at"`
}
