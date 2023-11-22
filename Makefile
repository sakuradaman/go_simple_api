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
	migrate -path build/db/migration/ddl/ -database 'mysql://root:@tcp(localhost:3306)/dramas?parseTime=true&loc=Local' up
# TODO: 以下が動かない所から確認
start:
	docker-compose exec app go run /cmd/main.go