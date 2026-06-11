package migrations

import (
	"context"
	"fmt"

	"gotodo/internal/model"

	"github.com/uptrace/bun"
)

func init() {
	// Up: 適用するときの処理
	up := func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*model.Todo)(nil)).
			Exec(ctx)
		if err != nil {
			return err
		}
		fmt.Println("created table: todos")
		return nil
	}

	// Down: 取り消すときの処理
	down := func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*model.Todo)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		fmt.Println("dropped table: todos")
		return nil
	}

	Migrations.MustRegister(up, down)
}