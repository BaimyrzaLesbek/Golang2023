package main

import (
	"Project/internal/data"
	"Project/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Manufacturer      string                 `json:"manufacturer"`
		StorageCapacity   int32                  `json:"storage_capacity"`
		Location          string                 `json:"location"`
		Resolution        string                 `json:"resolution"`
		FieldOfView       float32                `json:"field_of_view"`
		RecordingDuration data.RecordingDuration `json:"recording_duration"`
		PowerSource       string                 `json:"power_source"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	securityCamera := &data.SecurityCamera{
		Manufacturer:      input.Manufacturer,
		StorageCapacity:   input.StorageCapacity,
		Location:          input.Location,
		Resolution:        input.Resolution,
		FieldOfView:       input.FieldOfView,
		RecordingDuration: input.RecordingDuration,
		PowerSource:       input.PowerSource,
	}

	v := validator.New()

	if data.ValidateSecurityCamera(v, securityCamera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.SecurityCameras.Insert(securityCamera)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/security_cameras/%d", securityCamera.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": securityCamera}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showSecurityCamerasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	securityCamera, err := app.models.SecurityCameras.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"security_camera": securityCamera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
