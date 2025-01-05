package domain

import "time"

type Msg struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	RecipientId int64  `json:"recipient_id"`
	Message     string `json:"message"`
	CreatedAt   int64  `json:"created_at"`
}

type MsgOptions = func(*Msg)

func WithId(id int64) MsgOptions {
	return func(msg *Msg) {
		msg.Id = id
	}
}

func WithUserId(userId int64) MsgOptions {
	return func(msg *Msg) {
		msg.UserId = userId
	}
}

func WithRecipientId(recipientId int64) MsgOptions {
	return func(msg *Msg) {
		msg.RecipientId = recipientId
	}
}

func WithMessage(message string) MsgOptions {
	return func(msg *Msg) {
		msg.Message = message
	}
}

func WithCreatedAt(createdAt int64) MsgOptions {
	return func(msg *Msg) {
		msg.CreatedAt = createdAt
	}
}

func NewMsg(opts ...MsgOptions) *Msg {
	msg := &Msg{}

	msg.CreatedAt = time.Now().Unix()

	for _, fn := range opts {
		fn(msg)
	}

	return msg
}
