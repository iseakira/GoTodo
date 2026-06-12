package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"gotodo/internal/model"
	"gotodo/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type TodoHandler struct {
	repo *repository.TodoRepository
}

//構造体の実態をDB渡して生成する関数
func NewTodoHandler(db *bun.DB) *TodoHandler {
	return &TodoHandler{repo: repository.NewTodoRepository(db) }
}

//リクエストボディ用の型

type createTodoRequest struct {
	Content string `json:"content"`
	Until *time.Time `json:"until"`
}

type updateTodoRequest struct {
	Content string `json:"content"`
	Done *bool `json:"done"`
	Until *time.Time `json:"until"`
}

//GET /todos
func (h *TodoHandler) Index(c echo.Context) error {
	todos,err := h.repo.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err.Error()})
	}
	return c.JSON(http.StatusOK,todos)
}

//GET /todos/:id
func (h *TodoHandler) Show(c echo.Context) error {
	id,err := strconv.ParseInt(c.Param("id"),10,64)

	if err != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":"invalid id"})
	}

	todo,err := h.repo.Get(c.Request().Context(),id)

	if err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound,echo.Map{"error":"todo not found"})
		}
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err.Error()})
	}
	return c.JSON(http.StatusOK,todo)
}

//POST /todos
func (h *TodoHandler) Create(c echo.Context) error {
	var req createTodoRequest

	//.Bindでcにreqを埋め込み
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":"invalid request body"})
	}
	if req.Content == "" {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":"content is required"})
	}

	todo := &model.Todo{Content:req.Content}

	if req.Until != nil {
		todo.Until = *req.Until
	}

	if err := h.repo.Create(c.Request().Context(),todo); err != nil {
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err.Error()})
	}
	return c.JSON(http.StatusCreated,todo)

}

// PUT /todos/:id
func (h *TodoHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	var req updateTodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	// まず既存を取得
	todo, err := h.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// 送られてきた項目だけ書き換え
	if req.Content != "" {
		todo.Content = req.Content
	}
	if req.Done != nil {
		todo.Done = *req.Done
	}
	if req.Until != nil {
		todo.Until = *req.Until
	}
	todo.UpdatedAt = time.Now()

	if err := h.repo.Update(ctx, todo); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id
func (h *TodoHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	if err := h.repo.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}


