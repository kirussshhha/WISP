install:
	go mod download

run:
	nodemon -x "go run cmd/server/main.go" --signal SIGTERM -e go --verbose

lint:
	golangci-lint run