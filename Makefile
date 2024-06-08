build:
	go build -o ./bin/empl .
run: 
	docker compose build
	docker compose up
lint:
	golangci-lint run ./...
test:
	docker compose -f tests/docker-compose.yml build --no-cache
	docker compose -f tests/docker-compose.yml up 
