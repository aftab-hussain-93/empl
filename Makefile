build:
	go build -o ./bin/empl .
run: build
	./bin/empl
lint:
	golangci-lint run ./...