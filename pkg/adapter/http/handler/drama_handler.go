package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sakuradaman/go_simple_api/pkg/usecase"
)

type DramaHandler interface {
	SelectAllDramas() echo.HandlerFunc
}

type dramaHandler struct {
	usecase usecase.DramaUsecase
}

// usecaseインターフェースからhandlerインターフェースを作成
func NewDramaHandler(du usecase.DramaUsecase) *dramaHandler {
	return &dramaHandler{
		usecase: du,
	}
}

// /dramaに対するGETリクエストを処理
func (dh *dramaHandler) SelectAllDramas() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		drama, err := dh.usecase.SelectAllDramas(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, drama)
	}
}
