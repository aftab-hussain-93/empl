build:
	go build -o ./bin/empl .
run: 
	docker compose build --no-cache
	docker compose up
lint:
	golangci-lint run ./...
test_unit:
	go test -v ./...
test_e2e:
	docker compose -f tests/docker-compose.yml build
	docker compose -f tests/docker-compose.yml up 
