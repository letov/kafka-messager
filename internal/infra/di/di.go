package di

import (
	"kafka-messager/internal/infra/config"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/logger"
	"kafka-messager/internal/infra/msg"
	"kafka-messager/internal/infra/repo"

	"go.uber.org/fx"
)

func GetAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,

		db.NewDB,
		repo.NewBanWordDBRepo,
		repo.NewBlockUserDBRepo,

		msg.NewSchema,
		msg.NewEmitter,
		msg.NewReceiver,
	}
}

func InjectApp() fx.Option {
	return fx.Provide(
		GetAppConstructors()...,
	)
}
