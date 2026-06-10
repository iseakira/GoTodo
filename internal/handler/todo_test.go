package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestIndex(t *testing.T) {
	h := NewTodoHandler(nil)

	req := httptest.NewRequest(http.MethodGet,"/todos",nil)
	rec := httptest.NewRecorder()
	//リクエスト一件分の情報がコンテキスト
	c := echo.New().NewContext(req,rec)

	err := h.Index(c)

	if err != nil {
		t.Fatalf("Indexがエラーを返した: %v",err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("ステータスコード:200のはずが %d",rec.Code)
	}


}