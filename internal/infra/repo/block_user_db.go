package repo

import (
	"context"
	"kafka-messager/internal/infra/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type BlockUserDBRepo struct {
	pool *pgxpool.Pool
	log  *zap.SugaredLogger
}

func (bw *BlockUserDBRepo) Save(ctx context.Context, recipientId int64, blockUserId int64) error {
	query := `INSERT INTO public.blocked_users (recipient_id, block_user_id) VALUES (@recipient_id, @block_user_id)`
	args := pgx.NamedArgs{
		"recipient_id":  recipientId,
		"block_user_id": blockUserId,
	}
	_, err := bw.pool.Exec(ctx, query, args)
	return err
}

func NewBlockUserDBRepo(db *db.DB, log *zap.SugaredLogger) BlockedUser {
	return &BlockUserDBRepo{
		pool: db.GetPool(),
		log:  log,
	}
}
