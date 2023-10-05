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
		http.NotFound(w, r)
		return
	}

	securityCamera := data.SecurityCamera{
		ID:                id,
		CreatedAt:         time.Now(),
		StorageCapacity:   0,
		Location:          "",
		Resolution:        "",
		FieldOfView:       0,
		RecordingDuration: 0,
		PowerSource:       "",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"security_camera": securityCamera}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
