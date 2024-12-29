package emitter

import (
	"fmt"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/config"

	"github.com/lovoo/goka"
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

func NewEmitter(log *zap.SugaredLogger, config config.Config) *MsgEmitter {
	ge, err := goka.NewEmitter(config.Brokers, goka.Stream(config.MsgTopic), new(MsgCodec))
	if err != nil {
		log.Fatal(err)
	}
	return &MsgEmitter{
		ge,
		log,
	}
}
