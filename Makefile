docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

restart: down up

ksqldb-cli:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088

init: create-topics create-connectors ksqldb-migrations

create-topics:
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic postgres.public.blocked_users --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2 --config "cleanup.policy=compact"
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic postgres.public.ban_words --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2 --config "cleanup.policy=compact"
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic messages --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic filtered_messages --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2

create-connectors:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://localhost:8083/connectors/ -d @docker/postgres-connector.json

ksqldb-migrations:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088 -f '/docker/ksql-queries.sql'
