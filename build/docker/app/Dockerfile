# goのアプリケーションをビルドするためのバイナリファイルを作成
FROM golang:1.21.4-alpine as builder
# 作業ディレクトリを指定
WORKDIR /app
# ローカルの全てのファイルをコンテナのディレクトリにコピー
COPY . .
# 上記でコピーしたファイルを元にモジュールをコンテナにインストール
RUN go mod download
# goのアプリケーションをビルド、バイナリファイルをserverという名前で作成
RUN go build -o server cmd/main.go

# 上記で作成したバイナルファイルを実行するためのコンテナを作成
FROM alpine
# 正しいタイムゾーンで動かすためのパッケージをインストール
RUN apk --no-cache add tzdata
# 上記で作成したバイナリファイルをコンテナにコピー
COPY --from=builder /app/server server

ENTRYPOINT ["/server"]