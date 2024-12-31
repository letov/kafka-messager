package test

import (
	"context"
	"fmt"
	"kafka-messager/internal/infra/db"
	"kafka-messager/internal/infra/di"
	"math/rand"
	"os/exec"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func flushKafka(l *zap.SugaredLogger) error {
	cmd := exec.Command("make", "--file=../../Makefile", "init")
	out, err := cmd.Output()
	l.Info(string(out))
	return err
}

func flushDB(ctx context.Context, db *db.DB) error {
	pool := db.GetPool()
	query := `SELECT table_name "table" FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version';`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return err
	}

	var queries []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return err
		}
		queries = append(queries, fmt.Sprintf("TRUNCATE %v CASCADE;", table))
	}

	tx, _ := pool.Begin(ctx)
	for _, query := range queries {
		_, err = tx.Exec(ctx, query)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
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
