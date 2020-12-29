package main

import (
	"fmt"
	"go-http-training/api/users"
	"log"
	"net/http"
)

const (
	port = "8080"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	users.SetUserRoutes(srv)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)
	if err != nil {
		log.Panic(err)
	}
}
