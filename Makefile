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

create-topics: delete-topics
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic messages --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2

delete-topics:
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --if-exists --delete --topic postgres.public.blocked_users --bootstrap-server localhost:9093
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --if-exists --delete --topic messages --bootstrap-server localhost:9093

create-connectors: delete-connectors
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://localhost:8083/connectors/ -d @docker/postgres-connector.json

delete-connectors:
	curl -X DELETE http://localhost:8083/connectors/postgres-connector

ksqldb-migrations:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088 -f '/docker/ksql-queries.sql'

ksqldb-rollback:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088 -f '/docker/ksql-rollback.sql'
