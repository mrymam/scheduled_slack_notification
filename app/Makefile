
build:
	go build -o ./main ./main.go

exec: name=
exec:
	./main $(name)

run: name=
run:
	make build
	make exec name=$(name)

dev:
	GO_ENV=dev make run name=weekly_report_fourkeys

test:
	GO_ENV=test go test -v -cover -coverprofile=coverage.out ./...