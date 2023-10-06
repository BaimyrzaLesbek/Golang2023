package main

import (
	"Project/internal/data"
	"fmt"
	"net/http"
	"time"

	"encoding/json"
)

func (app *application) createSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Manufacturer    string  `json:"manufacturer"`
		StorageCapacity int32   `json:"storage_capacity"`
		Location        string  `json:"location"`
		Resolution      string  `json:"resolution"`
		FieldOfView     float32 `json:"field_of_view"`
		PowerSource     string  `json:"power_source"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
		Manufacturer:      "Panasonic",
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
