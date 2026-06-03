package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// main_test.go - サーバーの基本動作を確認するためのテスト
//
// このモジュールは、main.go で定義されたハンドラーが正しく動作することを検証します。

// TestHelloHandler は helloHandler 関数の動作をテストします。
func TestHelloHandler(t *testing.T) {
	// 擬似的なリクエストを作成
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// レスポンスを記録するための ResponseRecorder を作成
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// ハンドラーを実行
	handler.ServeHTTP(rr, req)

	// ステータスコードが 200 OK であることを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ハンドラーが返したステータスコードが正しくありません: 得られた値 %v、期待した値 %v",
			status, http.StatusOK)
	}

	// 期待される文字列が含まれていることを確認
	expectedStrings := []string{
		"Hello World",
		"pico.classless.min.css",
		"Go プロジェクトとしての基本的な構成",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("レスポンスボディに期待される文字列が含まれていません: %v", expected)
		}
	}
}
