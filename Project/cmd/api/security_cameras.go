package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new security camera")
}

func (app *application) showSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of security camera %d\n", id)
}
