package data

import (
	"Project/internal/validator"
	"time"
)

type SecurityCamera struct {
	ID                int64             `json:"id"`
	CreatedAt         time.Time         `json:"-"`
	Manufacturer      string            `json:"manufacturer"`
	StorageCapacity   int32             `json:"storage_capacity"`
	Location          string            `json:"location,omitempty"`
	Resolution        string            `json:"resolution"`
	FieldOfView       float32           `json:"field_of_view,string"`
	RecordingDuration RecordingDuration `json:"recording_duration"`
	PowerSource       string            `json:"power_source,omitempty"`
}

func ValidateSecurityCamera(v *validator.Validator, securityCamera *SecurityCamera) {
	v.Check(securityCamera.Manufacturer != "", "manufacturer", "must be provided")
	v.Check(len(securityCamera.Manufacturer) <= 500, "manufacturer", "must not be more than 500 bytes long")

	v.Check(securityCamera.StorageCapacity != 0, "storage_capacity", "must be provided")
	v.Check(securityCamera.StorageCapacity > 0, "storage_capacity", "must be positive")

	v.Check(securityCamera.Location != "", "location", "must be provided")
	v.Check(len(securityCamera.Location) <= 1000, "location", "must not be more than 1000 bytes long")

	v.Check(securityCamera.Resolution != "", "resolution", "must be provided")
	v.Check(len(securityCamera.Resolution) <= 50, "resolution", "must not be more than 50 bytes long")

	v.Check(securityCamera.FieldOfView != 0, "field_of_view", "must be provided")
	v.Check(securityCamera.FieldOfView > 0, "field_of_view", "must be positive")
	v.Check(securityCamera.FieldOfView <= 120, "field_of_view", "must not be more than 120 degrees")

	v.Check(securityCamera.RecordingDuration != 0, "recording_duration", "must be provided")
	v.Check(securityCamera.RecordingDuration > 0, "recording_duration", "must be positive")

	v.Check(securityCamera.PowerSource != "", "power_source", "must be provided")
	v.Check(len(securityCamera.PowerSource) <= 100, "power_source", "must not be more than 100 bytes")
}

//:param camera_id: Unique identifier for the camera (integer)
//:param created_at: Timestamp indicating when the camera was created (datetime)
//:param storage_capacity: Available storage capacity for recordings in gigabytes (float)
//:param location: Physical or virtual location of the camera (string)
//:param resolution: Resolution of the camera (e.g., 1080p, 4K) (string)
//:param field_of_view: The area covered by the camera (float)
//:param recording_duration: Maximum duration for recording in seconds (integer)
//:param power_source: Power source for the camera (e.g., wired, battery) (string)
