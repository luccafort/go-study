package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("wake up http server.")

	// リスト3.3設定を加えたWebサーバ
	server := http.Server{
		Addr:    "120.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
