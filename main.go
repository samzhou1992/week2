package main

import (
	"github.com/mytest/internal/handler"
	"log"
	"net/http"
)

func RegisterHandlers() {
	log.Println("register handlers")
	http.HandleFunc("/", handler.HeaderHandler)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	})
}

func init() {
	RegisterHandlers()
}

func main() {
	log.Println("start to listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
