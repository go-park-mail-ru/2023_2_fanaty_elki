network:
	docker network create mynetwork

.PHONY: up
up: 
	docker-compose up -d

.PHONY: down
down: 
	docker-compose down
