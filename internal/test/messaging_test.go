package test

import (
	"context"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/msg"
	"kafka-messager/internal/infra/repo"
	"strings"
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
			bw repo.BanWord,
			r *msg.Receiver,
			proc *msg.Processor,
		) {
			ctx := context.Background()
			_ = FlushDB(ctx, db)

			// 2->1 & 3->1 blocked
			_ = bu.Save(ctx, 1, 2)
			_ = bu.Save(ctx, 1, 3)

			_ = bw.Save(ctx, "bad_word1")
			_ = bw.Save(ctx, "bad_word2")

			msgOutCh := make(chan *domain.Msg)
			defer close(msgOutCh)
			msgInCh := make(chan interface{})

			doneCh := make(chan struct{})
			defer close(doneCh)

			go e.Emit(doneCh, msgOutCh)

			proc.UpdateBanWords(ctx)
			go proc.Run(ctx)
			time.Sleep(5 * time.Second)

			go r.Receive(doneCh, msgInCh)

			// 2->1 blocked
			msgOutCh <- domain.NewMsg(
				domain.WithId(1),
				domain.WithUserId(2),
				domain.WithRecipientId(1),
				domain.WithMessage("bad_word1 bad_word2 MESSAGE"),
			)

			// 3->1 blocked
			msgOutCh <- domain.NewMsg(
				domain.WithId(2),
				domain.WithUserId(3),
				domain.WithRecipientId(1),
				domain.WithMessage("bad_word1 bad_word2 MESSAGE"),
			)

			// 4->1 success
			msgOutCh <- domain.NewMsg(
				domain.WithId(3),
				domain.WithUserId(4),
				domain.WithRecipientId(1),
				domain.WithMessage("bad_word1 bad_word2 MESSAGE"),
			)

			raw := <-msgInCh
			m := raw.(*domain.Msg)

			assert.Equal(t, int64(4), m.UserId)
			assert.Equal(t, int64(1), m.RecipientId)
			assert.Equal(t, false, strings.Contains(m.Message, "bad_word1"))
			assert.Equal(t, false, strings.Contains(m.Message, "bad_word2"))

			proc.Stop()
			time.Sleep(5 * time.Second)

			doneCh <- struct{}{}
		})
	})
}
