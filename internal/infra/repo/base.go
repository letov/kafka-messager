package repo

import (
	"context"
	"kafka-messager/internal/domain"
)

type BanWord interface {
	GetList(ctx context.Context) ([]domain.BanWord, error)
	Save(ctx context.Context, word string) error
}

type BlockedUser interface {
	Save(ctx context.Context, recipientId int64, blockUserId int64) error
}
