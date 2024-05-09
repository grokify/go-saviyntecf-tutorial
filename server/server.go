package server

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	saviyntecf "github.com/grokify/go-saviyntecf"
)

type ECFAPI struct {
	saviyntecf.Unimplemented
}

func NewHandler() http.Handler {
	return saviyntecf.HandlerFromMux(NewECFAPI(), chi.NewRouter())

}

func NewECFAPI() saviyntecf.ServerInterface {
	return &ECFAPI{}
}
