package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"

	"gotodo/internal/config"
	"gotodo/internal/handler"
	"gotodo/internal/model"
)

func main() {
	config.LoadEnv()

	//DBを作成
	db := config.NewDB()
	defer db.Close()

	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*model.Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

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
