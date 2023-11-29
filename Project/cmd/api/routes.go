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

	router.HandlerFunc(http.MethodGet, "/v1/security_cameras", app.requireActivatedUser(app.listSec_CamerasHandler))
	router.HandlerFunc(http.MethodPost, "/v1/security_cameras", app.requireActivatedUser(app.createSecurityCamerasHandler))
	router.HandlerFunc(http.MethodGet, "/v1/security_cameras/:id", app.requireActivatedUser(app.showSecurityCamerasHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/security_cameras/:id", app.requireActivatedUser(app.updateSecurityCameraHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/security_cameras/:id", app.requireActivatedUser(app.deleteSecurityCameraHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
