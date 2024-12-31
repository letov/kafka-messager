package repo

import (
	"context"
)

type BanWord interface {
	Save(ctx context.Context, word string) error
}

type BlockedUser interface {
	Save(ctx context.Context, recipientId int64, blockUserId int64) error
}
