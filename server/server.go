package server

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	saviyntecf "github.com/grokify/go-saviyntecf"
)

func NewHandler() http.Handler {
	/*
		openapi, err := api.GetSwagger()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
			os.Exit(1)
		}

		openapi.Servers = nil
	*/

	// var service services.CareRequestService // TODO: in a production app, this would be properly initialised

	svr := NewECFAPI()

	r := chi.NewRouter()

	// NOTE that this is important! This enforces consumers use the right contract, required on top of checks that are done in the generated API code
	// r.Use(middleware.OapiRequestValidator(openapi))

	return saviyntecf.HandlerFromMux(svr, r)
}

func NewECFAPI() saviyntecf.ServerInterface {
	return &ECFAPI{}
}
