build:
	go build -o ./bin/empl .
run: 
	docker compose build
	docker compose up
lint:
	golangci-lint run ./...
