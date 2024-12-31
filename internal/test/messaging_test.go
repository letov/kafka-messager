package test

import (
	"context"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/kafka"
	"kafka-messager/internal/infra/repo"
	"testing"
	"time"

	"go.uber.org/zap"
)

func Test_BlockUser(t *testing.T) {
	t.Run("block user test", func(t *testing.T) {
		initTest(t, func(
			db *db.DB,
			e *kafka.MsgEmitter,
			l *zap.SugaredLogger,
			bu repo.BlockedUser,
		) {
			//_ = flushKafka(l)
			ctx := context.Background()
			_ = flushDB(ctx, db)

			// 2->1 & 3->1 blocked
			_ = bu.Save(ctx, 1, 2)
			_ = bu.Save(ctx, 1, 3)

			msgOutCh := make(chan *domain.Msg)
			defer close(msgOutCh)

			doneCh := make(chan struct{})
			defer close(doneCh)

			go e.Emit(doneCh, msgOutCh)

			// 2->1 blocked
			msg := domain.Msg{
				Id:          1,
				UserId:      2,
				RecipientId: 1,
				Message:     "msg",
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg

			// 3->1 blocked
			msg = domain.Msg{
				Id:          2,
				UserId:      3,
				RecipientId: 1,
				Message:     "msg",
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg

			// 4->1 success
			msg = domain.Msg{
				Id:          3,
				UserId:      4,
				RecipientId: 1,
				Message:     "msg",
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg
		})
	})
}
