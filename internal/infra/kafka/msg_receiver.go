package kafka

//
//import (
//	"kafka-messager/internal/infra/config"
//
//	"go.uber.org/fx"
//	"go.uber.org/zap"
//)
//
//type OrderReceiver struct {
//	cons       *kafka.Consumer
//	conf       *config.Config
//	l          *zap.SugaredLogger
//	sch        *Schema
//}
//
//func (or OrderReceiver) ReceiveMsg(
//	doneCh chan struct{},
//	outCh chan interface{},
//) {
//	go func() {
//		for {
//			select {
//			case <-doneCh:
//				close(outCh)
//				return
//			default:
//				ev := or.cons.Poll(pullTimeoutMs)
//				if ev == nil {
//					continue
//				}
//				switch e := ev.(type) {
//				case *kafka.Message:
//					or.l.Info("Get message")
//					data := dto.Order{}
//					err := or.sch.DeserializeInto(or.conf.OrdersTopic, e.Value, &data)
//					if err != nil {
//						or.l.Warn(err.Error())
//					} else {
//						if !or.autoCommit {
//							_, _ = or.cons.Commit()
//						}
//						outCh <- data
//					}
//				case kafka.Error:
//					or.l.Warn("Error: ", e)
//				default:
//					or.l.Warn("Some event: ", e)
//				}
//			}
//		}
//	}()
//}
//
//func NewOrderReceiver (
//	lc fx.Lifecycle,
//	conf *config.Config,
//	l *zap.SugaredLogger,
//	sch *Schema,
//) *OrderReceiver {
//	cons, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers":  conf.Brokers,
//		"group.id":           "consumer_group_1",
//		"session.timeout.ms": "6000",
//		"auto.offset.reset":  "earliest",
//		"enable.auto.commit": "true",
//		"acks":               "all",
//	})
//	if err != nil {
//		orf.l.Fatal("Error creating consumer: ", err)
//	}
//
//	orf.l.Info("Consumer created")
//
//	err = cons.SubscribeTopics([]string{orf.conf.OrdersTopic}, nil)
//	if err != nil {
//		orf.l.Fatal("Subscribe error:", err)
//	}
//
//	orf.cs = append(orf.cs, cons)
//	return &OrderReceiver{cons, orf.conf, orf.l, orf.sch, autoCommit}
//}
