package test

import (
	"context"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/msg"
	"kafka-messager/internal/infra/repo"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"go.uber.org/zap"
)

func Test_BlockUser(t *testing.T) {
	t.Run("block user test", func(t *testing.T) {
		initTest(t, func(
			db *db.DB,
			e *msg.MsgEmitter,
			l *zap.SugaredLogger,
			bu repo.BlockedUser,
			r *msg.Receiver,
		) {
			ctx := context.Background()
			_ = flushDB(ctx, db)

			// 2->1 & 3->1 blocked
			_ = bu.Save(ctx, 1, 2)
			_ = bu.Save(ctx, 1, 3)

			msgOutCh := make(chan *domain.Msg)
			defer close(msgOutCh)
			msgInCh := make(chan interface{})

			doneCh := make(chan struct{})
			defer close(doneCh)

			go e.Emit(doneCh, msgOutCh)
			go r.Receive(doneCh, msgInCh)

			// 2->1 blocked
			msg1 := domain.Msg{
				Id:          1,
				UserId:      2,
				RecipientId: 1,
				Message:     RandString(10),
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg1

			// 3->1 blocked
			msg2 := domain.Msg{
				Id:          2,
				UserId:      3,
				RecipientId: 1,
				Message:     RandString(10),
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg2

			// 4->1 success
			msg3 := domain.Msg{
				Id:          3,
				UserId:      4,
				RecipientId: 1,
				Message:     RandString(10),
				CreatedAt:   time.Now().Unix(),
			}
			msgOutCh <- &msg3

			raw := <-msgInCh
			m := raw.(*domain.Msg)

			assert.Equal(t, int64(4), m.UserId)
			assert.Equal(t, int64(1), m.RecipientId)

			doneCh <- struct{}{}
		})
	})
}
