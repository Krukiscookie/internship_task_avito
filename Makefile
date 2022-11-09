.SILENT:

build:
	docker-compose build intern_task

run: build
	docker-compose up intern_task

stop:
	docker-compose down --remove-orphans

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:test@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go

