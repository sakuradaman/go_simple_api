init:
	@make build
	@make up
build:
# docker-compose記載の内容でdockerコンテナイメージをビルドする
	docker-compose build --no-cache
up:
# 上記で作成したイメージを元にコンテナを起動する
	docker-compose up -d
migrate:
	docker-compose exec app -dir /db/migration mysql "root:user@tcp(mysql:3306)/dramas" up
start:
	docker-compose exec app go run /cmd/main.go