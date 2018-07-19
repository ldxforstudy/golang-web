package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome"))
	})
	http.ListenAndServe(":3000", r)
}
