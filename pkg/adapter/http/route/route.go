package route

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sakuradaman/go_simple_api/pkg/lib/config"
	"golang.org/x/xerrors"
)

type Route interface {
	InitRouting(*config.Config) (*echo.Echo, error)
}

type InitRoute struct {
}

func NewInitRoute() Route {
	InitRoute := InitRoute{}
	return &InitRoute
}

func (i *InitRoute) InitRouting(cfg *config.Config) (*echo.Echo, error) {
	// インスタンスの生成
	e := echo.New()

	e.Use(
		// httpリクエストとレスポンスのログを出力する
		middleware.Logger(),
		// リクエストの復元
		middleware.Recover(),
	)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// エラーをログに出力
		e.Logger.Error(err)

		// エラーをJSONで返す
		if he, ok := err.(*echo.HTTPError); ok {
			c.JSON(he.Code, he.Message)
		} else {
			c.JSON(500, "Internal Server Error")
		}
	}

	// ルーティングの設定
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.GET("/drama")

	// Start Server
	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		return nil, xerrors.Errorf("fail to start port:%s %w", cfg.Port, err)
	}

	return e, nil
}