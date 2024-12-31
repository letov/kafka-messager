package repo

import (
	"context"
	"kafka-messager/internal/infra/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type BanWordDBRepo struct {
	pool *pgxpool.Pool
	log  *zap.SugaredLogger
}

func (bw *BanWordDBRepo) Save(ctx context.Context, word string) error {
	query := `INSERT INTO public.ban_words (word) VALUES (@word)`
	args := pgx.NamedArgs{
		"word": word,
	}
	_, err := bw.pool.Exec(ctx, query, args)
	return err
}

func NewBanWordDBRepo(db *db.DB, log *zap.SugaredLogger) BanWord {
	return &BanWordDBRepo{
		pool: db.GetPool(),
		log:  log,
	}
}
