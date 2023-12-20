network:
	docker network create mynetwork

.PHONY: up
up: 
	docker-compose build
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
	
.PHONY: copy
copy:
	mkdir -p AuthService/build && cp build/auth AuthService/build
	mkdir -p ProductService/build && cp build/product ProductService/build
	mkdir -p UserService/build && cp build/user UserService/build

.PHONY: test
test:
	cd server && go test -coverprofile=c.out ./... && go tool cover -func c.out | grep total