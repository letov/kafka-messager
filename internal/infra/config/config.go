package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDns                string
	MsgTopic                   string
	MsgFilteredBlockUsersTopic string
	SchemaRegistryUrl          string
	Brokers                    []string
	MsgFiltered                string
	KafkaSessionTimeoutMs      int
	KafkaAutoOffsetReset       string
	KafkaConsumerPullTimeoutMs int
	KafkaAcks                  string
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
		DatabaseDns:                getEnvStr("DATABASE_DSN", ""),
		MsgTopic:                   getEnvStr("MSG_TOPIC", ""),
		MsgFilteredBlockUsersTopic: getEnvStr("MSG_FILTERED_BLOCK_USERS_TOPIC", ""),
		SchemaRegistryUrl:          getEnvStr("SCHEMA_REGISTRY_URL", ""),
		Brokers:                    strings.Split(getEnvStr("KAFKA_BROKERS", ""), ","),
		MsgFiltered:                getEnvStr("MSG_FILTERED_TOPIC", ""),
		KafkaSessionTimeoutMs:      getEnvInt("KAFKA_SESSION_TIMEOUT_MS", 0),
		KafkaAutoOffsetReset:       getEnvStr("KAFKA_AUTO_OFFSET_RESET", ""),
		KafkaConsumerPullTimeoutMs: getEnvInt("KAFKA_CONSUMER_PULL_TIMEOUT_MS", 0),
		KafkaAcks:                  getEnvStr("KAFKA_ACKS", ""),
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
