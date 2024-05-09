package main

import (
	"log"
	"net/http"
	"time"

	"github.com/grokify/go-saviyntecf-tutorial/server"
)

func main() {
	timeout := 1 * time.Second
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           server.NewHandler(),
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
