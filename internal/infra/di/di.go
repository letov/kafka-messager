package di

import (
	"kafka-messager/internal/infra/config"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/logger"

	"go.uber.org/fx"
)

func GetConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,
		db.NewDB,
	}
}

func InjectApp() fx.Option {
	return fx.Provide(
		GetConstructors()...,
	)
}
