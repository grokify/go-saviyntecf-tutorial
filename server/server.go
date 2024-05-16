package server

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	saviyntecf "github.com/grokify/go-saviyntecf"
)

const (
	AccountsImportCSVURL = "https://raw.githubusercontent.com/grokify/go-saviyntecf-tutorial/main/ECF_Account_Import.csv"
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
