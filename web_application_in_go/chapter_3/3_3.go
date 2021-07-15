package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// https://github.com/golang/go/blob/21a04e33353316635b5f3351e807916f3bb1e844/src/net/http/server.go#L86-L88
// ↑のコードでinterfaceが定義されている。
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
