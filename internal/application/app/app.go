package app

import (
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/msg"
)

func Start(
	_ *db.DB,
	_ *msg.MsgEmitter,
	_ *msg.Receiver,
	_ *msg.Processor,
) {
}
