docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

ksqldb-cli:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088

create-connectors:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://localhost:8083/connectors/ -d @postgres-connector.json

create-topics:
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic user-likes-stream --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics  --describe --topic user-likes-stream  --bootstrap-server localhost:9092
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic user-like-group-table --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2 --config cleanup.policy=compact
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics  --describe --topic user-like-group-table  --bootstrap-server localhost:9092