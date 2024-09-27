package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/results", app.listResultsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/results", app.createResultHandler)
	router.HandlerFunc(http.MethodGet, "/v1/results/:id", app.showResultHandler)

	return app.recoverPanic(router)
}
