package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/security_cameras", app.listSec_CamerasHandler)
	router.HandlerFunc(http.MethodPost, "/v1/security_cameras", app.createSecurityCamerasHandler)
	router.HandlerFunc(http.MethodGet, "/v1/security_cameras/:id", app.showSecurityCamerasHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/security_cameras/:id", app.updateSecurityCameraHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/security_cameras/:id", app.deleteSecurityCameraHandler)

	return app.recoverPanic(app.rateLimit(router))
}
