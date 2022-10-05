
build:
	go build -o ./main ./main.go

exec:
	./main

run:
	make build
	make exec

cp-files:
	cp ../setting.yaml internal/setting/files/setting.yaml
	cp ../query/sample.sql internal/query/files/sample.sql
	cp ../templates/sample.md internal/view/templates/sample.md

dev:
	make cp-files
	docker-compose run --rm \
	-e GO_ENV=dev dev \
	make run

test:
	make cp-files
	docker-compose run --rm \
	-e GO_ENV=test dev \
	go test -v -cover -coverprofile=coverage.out ./...