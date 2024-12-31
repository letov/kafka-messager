package di

import (
	"kafka-messager/internal/infra/config"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/kafka"
	"kafka-messager/internal/infra/logger"
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

		kafka.NewSchema,
		kafka.NewEmitter,
	}
}

func InjectApp() fx.Option {
	return fx.Provide(
		GetAppConstructors()...,
	)
}
