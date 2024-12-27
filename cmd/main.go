package main

import (
	"kafka-messager/internal/app/app"
	"kafka-messager/internal/infra/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(app.Start),
	).Run()
}
