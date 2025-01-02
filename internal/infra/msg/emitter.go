package msg

import (
	"context"
	"fmt"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/config"

	"github.com/lovoo/goka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MsgEmitter struct {
	ge  *goka.Emitter
	log *zap.SugaredLogger
}

func (me MsgEmitter) Emit(doneCh chan struct{}, msgCh chan *domain.Msg) {
	for {
		select {
		case <-doneCh:
			return
		case msg := <-msgCh:
			err := me.ge.EmitSync(fmt.Sprintf("msg-%d", msg.Id), msg)
			if err != nil {
				me.log.Fatal(err)
			}
		}
	}
}

func NewEmitter(lc fx.Lifecycle, log *zap.SugaredLogger, config config.Config, sch *Schema) *MsgEmitter {
	codec := NewMsgCodec(config.MsgTopic, sch)
	ge, err := goka.NewEmitter(config.Brokers, goka.Stream(config.MsgTopic), codec)
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return ge.Finish()
		},
	})

	return &MsgEmitter{
		ge,
		log,
	}
}
