package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/security_cameras", app.createSecurityCamerasHandler)
	router.HandlerFunc(http.MethodGet, "/v1/security_cameras/:id", app.showSecurityCamerasHandler)

	return router
}
