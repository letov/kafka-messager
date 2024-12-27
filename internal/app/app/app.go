package app

import (
	"context"
	"kafka-messager/internal/infra/db"
)

func Start(
	db *db.DB,
) {
	db.Ping(context.Background())
}
