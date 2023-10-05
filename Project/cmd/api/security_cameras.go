package main

import (
	"fmt"
	"net/http"
	"time"

	"Project/internal/data"
)

func (app *application) createSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new security camera")
}

func (app *application) showSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	securityCamera := data.SecurityCamera{
		ID:                id,
		CreatedAt:         time.Now(),
		StorageCapacity:   900,
		Location:          "Tole bi",
		Resolution:        "1080p",
		FieldOfView:       500,
		RecordingDuration: 100,
		PowerSource:       "wire",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"security_camera": securityCamera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
