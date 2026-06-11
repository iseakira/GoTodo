package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun/migrate"

	"gotodo/internal/config"
	"gotodo/migrations"
)

func main() {
	config.LoadEnv()

	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./cmd/migrate [init|up|down|status]")
	}
	cmd := os.Args[1]

	db := config.NewDB()
	defer db.Close()

	migrator := migrate.NewMigrator(db, migrations.Migrations)
	ctx := context.Background()

	switch cmd {
	case "init":
		if err := migrator.Init(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("migration tables created")

	case "up":
		group, err := migrator.Migrate(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if group.IsZero() {
			fmt.Println("no new migrations to run")
			return
		}
		fmt.Printf("migrated to %s\n", group)

	case "down":
		group, err := migrator.Rollback(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if group.IsZero() {
			fmt.Println("no migrations to roll back")
			return
		}
		fmt.Printf("rolled back %s\n", group)

	case "status":
		ms, err := migrator.MigrationsWithStatus(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("migrations: %s\n", ms)
		fmt.Printf("unapplied: %s\n", ms.Unapplied())
		fmt.Printf("last applied: %s\n", ms.LastGroup())

	default:
		log.Fatalf("unknown command: %s", cmd)
	}
}