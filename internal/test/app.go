package test

import (
	"kafka-messager/internal/infra/di"
	"math/rand"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetAppTestConstructors() []interface{} {
	return []interface{}{}
}

func InjectAppTest() fx.Option {
	return fx.Provide(
		append(
			di.GetAppConstructors(),
			GetAppTestConstructors()...,
		)...,
	)
}

func initTest(t *testing.T, r interface{}) {
	t.Setenv("IS_TEST_ENV", "true")
	app := fxtest.New(t, InjectAppTest(), fx.Invoke(r))
	defer app.RequireStop()
	app.RequireStart()
}
