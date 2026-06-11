package main

import (
	"github.com/labstack/echo/v4"

	"gotodo/internal/config"
	"gotodo/internal/handler"
)

func main() {
	config.LoadEnv()

	db := config.NewDB()
	defer db.Close()

	//DBを持ったhandlerある意味依存性の注入
	h := handler.NewTodoHandler(db)

	e := echo.New()
	e.GET("/todos", h.Index)
	e.GET("/todos/:id",h.Show)
	e.POST("/todos", h.Create)
	e.PUT("/todos/:id",h.Update)
	e.DELETE("/todos/:id", h.Delete)

	e.Logger.Fatal(e.Start(":8989"))
}
