package msg

import (
	"context"
	"encoding/binary"
	"errors"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/config"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Receiver struct {
	cons  *kafka.Consumer
	conf  config.Config
	l     *zap.SugaredLogger
	sch   *Schema
	codec *Codec
}

func (r Receiver) Receive(
	doneCh chan struct{},
	msgCh chan interface{},
) {
	go func() {
		for {
			select {
			case <-doneCh:
				close(msgCh)
				return
			default:
				ev := r.cons.Poll(r.conf.KafkaConsumerPullTimeoutMs)
				if ev == nil {
					continue
				}
				switch e := ev.(type) {
				case *kafka.Message:
					r.l.Info("Get message")
					raw, err := r.codec.Decode(e.Value)
					data, ok := raw.(*domain.Msg)
					if !ok {
						r.l.Warn(errors.New("decode message error"))
					}
					//err := r.sch.DeserializeInto(r.conf.MsgFilteredBlockUsersTopic, e.Value, &data)
					if err != nil {
						r.l.Warn(err.Error())
					} else {
						data.RecipientId = int64(binary.BigEndian.Uint32(e.Key))
						_, _ = r.cons.Commit()
						msgCh <- data
					}
				case kafka.Error:
					r.l.Warn("Error: ", e)
				default:
					r.l.Warn("Some event: ", e)
				}
			}
		}
	}()
}

func NewReceiver(
	lc fx.Lifecycle,
	conf config.Config,
	l *zap.SugaredLogger,
	sch *Schema,
) *Receiver {
	cons, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(conf.Brokers, ","),
		"group.id":           conf.MsgFiltered + "-for-test",
		"session.timeout.ms": conf.KafkaSessionTimeoutMs,
		"auto.offset.reset":  conf.KafkaAutoOffsetReset,
		"acks":               conf.KafkaAcks,
	})
	if err != nil {
		l.Fatal("Error creating consumer: ", err)
	}

	l.Info("Consumer created")

	if cons.SubscribeTopics([]string{conf.MsgFiltered + "-table"}, nil) != nil {
		l.Fatal("Subscribe error:", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return cons.Close()
		},
	})

	codec := new(Codec)
	return &Receiver{cons, conf, l, sch, codec}
}
