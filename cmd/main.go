package main

import (
	"fmt"

	"golang.org/x/exp/slog"

	"github.com/sakuradaman/go_simple_api/pkg/infra"
	"github.com/sakuradaman/go_simple_api/pkg/lib/config"
)

func main() {
	// Config（環境変数）の読み取り
	c, cErr := config.New()

	// Loggingの設定(必要であれば後でやる)

	// DBに接続
	db, err := infra.NewDBConnector(c)
	if err != nil {
		slog.Error("initDb err: %w", err)
	}
	// DBの構造体を定義
	dr := infra.NewDramaRepository(db)
	fmt.Println(dr, "debug")

	// TODO: drを使ったユースケース（今回作成するAPIの処理)を記載

	// URLのルーティングの設定
	r := route.NewInitRoute(dr)

	defer func() {
		// 指定した環境変数が存在しない場合
		if cErr != nil {
			slog.Error("config err: %w", cErr)
		}
	}()
}
