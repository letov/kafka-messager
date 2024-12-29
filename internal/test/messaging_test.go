package test

import (
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/emitter"
	"math/rand"
	"testing"
	"time"

	"go.uber.org/zap"
)

func Test_Messaging(t *testing.T) {
	t.Run("test messaging", func(t *testing.T) {
		initTest(t, func(
			e *emitter.MsgEmitter,
			l *zap.SugaredLogger,
		) {
			msgCnt := 500

			msgOutCh := make(chan *domain.Msg, 50)
			defer close(msgOutCh)

			doneCh := make(chan struct{})
			defer close(doneCh)

			for i := 0; i < 5; i++ {
				go e.Emit(doneCh, msgOutCh)
			}

			for i := 1; i <= msgCnt; i++ {
				msg := domain.Msg{
					Id:          int64(i),
					UserId:      int64(rand.Uint32()),
					RecipientId: int64(rand.Uint32()),
					Message:     RandString(50),
					Timestamp:   int64(rand.Uint64()),
				}
				msgOutCh <- &msg
				time.Sleep(time.Millisecond * 1)
			}
		})
	})
}
