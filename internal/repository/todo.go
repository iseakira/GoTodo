package repository

import (
	"context"
	"gotodo/internal/model"

	"github.com/uptrace/bun"
)

type TodoRepository struct {
	db *bun.DB
}

//一覧取得
func (r *TodoRepository) List(ctx context.Context) ([]model.Todo,error) {
	var todos []model.Todo
	err := r.db.NewSelect().Model(&todos).Order("id ASC").Scan(ctx)
	return todos,err
}