APP_NAME=crud

build: dep
	CGO_ENABLED=0 GOOS=linux go build -ldflags -a -installsuffix nocgo -o ./bin ./...

dep:
	@echo ">> Downloading Dependencies"
	@go mod download

run-api:
	@echo ">> Running API Server"
	@go run main.go serve-http

swag-init:
	@echo ">> Running swagger init"
	@swag init

run-consumer:
	@echo ">> Running Consumer"
	@go run main.go run-consumer

remock:
	#https://github.com/vektra/mockery

	@echo ">> Mock Domain"
	@mockery

run-test: dep
	@echo ">> Running Test"
	@go test -v -cover -count=1 -failfast -covermode=atomic ./...