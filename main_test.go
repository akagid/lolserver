package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	// テスト用のHTTPリクエストを作成
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// HTTPレスポンスを記録するためのレコーダーを作成
	// httptest.NewRecorder を使用することで実際のネットワークI/Oを省略し、リクエストとレスポンスをメモリ内でエミュレート可能
	recorder := httptest.NewRecorder()

	// request の内容をリクエストし、レスポンスを recorder に記録
	router.ServeHTTP(recorder, request)

	// ステータスコードのチェック
	assert.Equal(t, http.StatusOK, recorder.Code)

	// レスポンスボディのJSONチェック
	expected := `{"message":"pong"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}
