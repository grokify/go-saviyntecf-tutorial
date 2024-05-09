package main

import (
	"log"
	"net/http"

	"github.com/grokify/go-saviyntecf-tutorial/server"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.NewHandler(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
