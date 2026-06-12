package repository

import (
	"context"
	"gotodo/internal/model"

	"github.com/uptrace/bun"
)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository(db *bun.DB) *TodoRepository {
	return &TodoRepository{db:db}
}

//一覧取得
//SELECT * FROM todos ORDER BY id ASC;
//Model(&todos)でtodosに結果を入れることを指定
func (r *TodoRepository) List(ctx context.Context) ([]model.Todo,error) {
	var todos []model.Todo
	err := r.db.NewSelect().Model(&todos).Order("id ASC").Scan(ctx)
	return todos,err
}

//一件取得
func (r *TodoRepository) Get(ctx context.Context, id int64) (*model.Todo,error) {
	todo := new(model.Todo)
	err := r.db.NewSelect().Model(todo).Where("id=?",id).Scan(ctx)

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	_,err := r.db.NewInsert().Model(todo).Exec(ctx)

	return err
}

func (r *TodoRepository) Update(ctx context.Context,todo *model.Todo) error {
	_,err := r.db.NewUpdate().Model(todo).WherePK().OmitZero().Exec(ctx)

	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id int64) error {
	_,err := r.db.NewDelete().Model((*model.Todo)(nil)).Where("id=?",id).Exec(ctx)

	return err
}