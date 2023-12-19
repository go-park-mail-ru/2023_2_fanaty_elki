network:
	docker network create mynetwork

.PHONY: up
up: 
	docker-compose up -d

.PHONY: down
down: 
	docker-compose down

GOBUILD = go build -buildvcs=false

.PHONY: build
build:
	mkdir -p build
	cd server && ${GOBUILD} -o ../build/server cmd/server.go
	cd AuthService && ${GOBUILD}  -o ../build/auth cmd/main.go
	cd UserService && ${GOBUILD}  -o ../build/user cmd/main.go
	cd ProductService && ${GOBUILD}  -o ../build/product cmd/main.go