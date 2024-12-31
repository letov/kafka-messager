package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDns       string
	Brokers           []string
	MsgTopic          string
	SchemaregistryUrl string
}

func NewConfig() Config {
	var err error
	if os.Getenv("IS_TEST_ENV") == "true" {
		err = godotenv.Load("../../.env.test")
	} else {
		err = godotenv.Load(".env")
	}

	if err != nil {
		panic(err)
	}

	return Config{
		DatabaseDns:       getEnvStr("DATABASE_DSN", ""),
		Brokers:           strings.Split(getEnvStr("BROKERS", ""), ","),
		MsgTopic:          getEnvStr("MSG_TOPIC", ""),
		SchemaregistryUrl: getEnvStr("SCHEMA_REGISTRY_URL", ""),
	}
}

func getEnvInt(key string, def int) int {
	v, e := strconv.Atoi(getEnvStr(key, strconv.Itoa(def)))
	if e != nil {
		return def
	} else {
		return v
	}
}

func getEnvStr(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}
