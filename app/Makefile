
build:
	go build -o ./main ./main.go

exec:
	./main

run: build exec

test:
	GO_ENV=test go test -v -cover -coverprofile=coverage.out ./...