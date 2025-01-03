package repo

import (
	"context"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type BanWordDBRepo struct {
	pool *pgxpool.Pool
	log  *zap.SugaredLogger
}

func (bw *BanWordDBRepo) GetList(ctx context.Context) ([]domain.BanWord, error) {
	query := `SELECT * FROM public.ban_words`
	rows, err := bw.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var res []domain.BanWord
	for rows.Next() {
		var r domain.BanWord
		err = rows.Scan(
			&r.Id,
			&r.Word,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}

	return res, nil
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
