
build:
	go build -o ./main ./main.go

exec:
	./main

run: build exec

dev:
	GO_ENV=dev make run

test:
	GO_ENV=test go test -v -cover -coverprofile=coverage.out ./...