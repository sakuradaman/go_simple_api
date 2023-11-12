package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type dramaHandler struct {
	usecase usecase.DramaUsecase
}

func NewDramaHandler() *dramaHandler {
	return &dramaHandler{}
}

// /dramaに対するGETリクエストを処理
func (dh *dramaHandler) FindAllDramas() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		drama, err := dh.usecase.FindAllDramas(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, drama)
	}
}
