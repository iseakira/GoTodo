package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type TodoHandler struct {
	db *bun.DB
}

//構造体の実態をDB渡して生成する関数
func NewTodoHandler(db *bun.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

//handlerのメソッド
func (h *TodoHandler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
