
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
	docker-compose run --rm \
	-e GO_ENV=dev dev \
	make run schedule=weekly

test:
	docker-compose run --rm \
	-e GO_ENV=test dev \
	go test -v -cover -coverprofile=coverage.out ./...