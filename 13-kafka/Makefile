COMPOSE=docker compose
FILE=docker-compose.yml
BROKER=broker-1
BOOTSTRAP=broker-1:19092

.PHONY: up down reset topics consume

up:
	$(COMPOSE) -f $(FILE) up -d

down:
	$(COMPOSE) -f $(FILE) down

reset:
	$(COMPOSE) -f $(FILE) down -v --remove-orphans

topics:
	docker exec -it $(BROKER) bash -c "/opt/kafka/bin/kafka-topics.sh --bootstrap-server $(BOOTSTRAP) --list"

consume:
	docker exec -it broker-1 bash -c "/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server broker-1:19092 --topic $(t) --from-beginning --group debug1"