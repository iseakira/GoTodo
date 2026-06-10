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

	db := config.NewDB()
	defer db.Close()

	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*model.Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewTodoHandler(db)

	e := echo.New()
	e.GET("/", h.Index)
	e.Logger.Fatal(e.Start(":8989"))
}
