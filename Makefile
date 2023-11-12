init:
	@make build
	@make up
build:
	docker-compose build --no-cache
up:
	docker-compose up -d
migrate:
	docker-compose exec app goose -dir ./build/db/migration mysql "root:user@tcp(mysql:3306)/golang" up
start:
	docker-compose exec app go run ./cmd/main.go