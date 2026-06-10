package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"gotodo/internal/config"
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

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
