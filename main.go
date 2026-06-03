package main

import (
	"fmt"
	"net/http"
)

// main.go - アプリケーションのメインエントリーポイント
//
// このモジュールは、net/http パッケージを使用して基本的な HTTP サーバーを起動します。
// ルートパスへのアクセスに対して、Pico CSS を適用した Hello World ページを返します。

// helloHandler は、ルートパス ("/") への HTTP リクエストを処理します。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 日本語を正しく表示するために Content-Type を設定
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Pico CSS (classless バージョン) を適用した HTML を出力
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.classless.min.css">
    <title>Hello World</title>
</head>
<body>
    <main>
        <h1>Hello World</h1>
        <p>Go プロジェクトとしての基本的な構成（Step 1）が完了しました。</p>
        <p>この画面は net/http と Pico CSS を使用して表示されています。</p>
    </main>
</body>
</html>
`)
}

// main はアプリケーションの開始点です。
// サーバーのルーティングを設定し、ポート 8080 で待機を開始します。
func main() {
	// ルートパスのハンドラーを登録
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on http://localhost:8080...")

	// 8080ポートでサーバーを起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %s\n", err)
	}
}
