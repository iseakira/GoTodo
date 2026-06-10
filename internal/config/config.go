package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func LoadEnv() {
	godotenv.Load()
}

func NewDB() *bun.DB {
	sqldb, err := sql.Open("pg", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
