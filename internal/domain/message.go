package domain

type Msg struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	RecipientId int64  `json:"recipient_id"`
	Message     string `json:"message"`
	CreatedAt   int64  `json:"created_at"`
}
