package main

import (
	"net/http"
)

func main() {
	personMux := http.NewServeMux()
	personMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, web!"))
		})
	dogMux := http.NewServeMux()
	dogMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, dog!"))
		})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", personMux))
	mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
