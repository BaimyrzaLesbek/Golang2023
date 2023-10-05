package data

import "time"

type SecurityCamera struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"-"`
	StorageCapacity   int32     `json:"storage_capacity"`
	Location          string    `json:"location,omitempty"`
	Resolution        string    `json:"resolution"`
	FieldOfView       float32   `json:"field_of_view,string"`
	RecordingDuration int64     `json:"recording_duration"`
	PowerSource       string    `json:"power_source,omitempty"`
}

//:param camera_id: Unique identifier for the camera (integer)
//:param created_at: Timestamp indicating when the camera was created (datetime)
//:param storage_capacity: Available storage capacity for recordings in gigabytes (float)
//:param location: Physical or virtual location of the camera (string)
//:param resolution: Resolution of the camera (e.g., 1080p, 4K) (string)
//:param field_of_view: The area covered by the camera (float)
//:param recording_duration: Maximum duration for recording in seconds (integer)
//:param power_source: Power source for the camera (e.g., wired, battery) (string)
