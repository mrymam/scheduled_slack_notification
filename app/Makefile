
build:
	go build -o ./main ./main.go

exec: schedule=
exec:
	./main $(schedule)

run: schedule=
run:
	make build
	make exec schedule=$(schedule)

dev:
	GO_ENV=dev make run schedule=weekly

test:
	GO_ENV=test go test -v -cover -coverprofile=coverage.out ./...