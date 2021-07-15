package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// https://github.com/golang/go/blob/21a04e33353316635b5f3351e807916f3bb1e844/src/net/http/server.go#L2485-L2490
	// ↑で定義されている関数。
	// 第2引数で関数を受け取るがその際に受け取る関数の引数で型が指定されているため、
	// 登録される関数の動作が保証されている
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
