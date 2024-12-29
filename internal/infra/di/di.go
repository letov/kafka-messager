package di

import (
	"kafka-messager/internal/infra/config"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/emitter"
	"kafka-messager/internal/infra/logger"

	"go.uber.org/fx"
)

func GetAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,
		db.NewDB,
		emitter.NewEmitter,
	}
}

func InjectApp() fx.Option {
	return fx.Provide(
		GetAppConstructors()...,
	)
}
