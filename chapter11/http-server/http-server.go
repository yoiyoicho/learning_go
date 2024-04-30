package main

import (
	"fmt"
	"net/http"
	"time"
)

// インタフェースhttp.Handlerはひとつだけメソッドを持つ
// type Handler interface {
// 	ServeHTTP(http.ResponseWriter, *http/Request)
// }

// http.ResponseWriterは3つのメソッドを持つ
// type ResponseWriter interface {
// 	Header() http.Header
// 	Write([]byte) (int, error)
// 	WriteHeader(statusCode int)
// }

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}
	fmt.Println("Starting server on :8080")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
