package data

import (
	"Project/internal/validator"
	"context"
	"database/sql"
	"errors"
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
	Version           int32             `json:"version"`
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

type SecurityCameraModel struct {
	DB *sql.DB
}

func (s SecurityCameraModel) Insert(SecurityCamera *SecurityCamera) error {
	query := `
	INSERT INTO security_cameras (manufacturer, storage_capacity, location, resolution, field_of_view, recording_duration, power_source)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id, created_at, version
	`
	args := []interface{}{SecurityCamera.Manufacturer, SecurityCamera.StorageCapacity, SecurityCamera.Location, SecurityCamera.Resolution, SecurityCamera.FieldOfView, SecurityCamera.RecordingDuration, SecurityCamera.PowerSource}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&SecurityCamera.ID, &SecurityCamera.CreatedAt, &SecurityCamera.Version)
}

func (s SecurityCameraModel) Get(id int64) (*SecurityCamera, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT id, created_at, manufacturer, storage_capacity, location, resolution, field_of_view, recording_duration, power_source, version 
	FROM security_cameras 
	WHERE id = $1
	`
	var security_camera SecurityCamera

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&security_camera.ID,
		&security_camera.CreatedAt,
		&security_camera.Manufacturer,
		&security_camera.StorageCapacity,
		&security_camera.Location,
		&security_camera.Resolution,
		&security_camera.FieldOfView,
		&security_camera.RecordingDuration,
		&security_camera.PowerSource,
		&security_camera.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &security_camera, nil
}

func (s SecurityCameraModel) Update(SecurityCamera *SecurityCamera) error {
	query := `
	UPDATE security_cameras SET manufacturer = $1, storage_capacity = $2, location = $3, resolution = $4,
	                            field_of_view = $5, recording_duration = $6, power_source = $7, version = version + 1
	                        WHERE id = $8 and version = $9
	                        RETURNING version;
	`
	args := []interface{}{
		SecurityCamera.Manufacturer,
		SecurityCamera.StorageCapacity,
		SecurityCamera.Location,
		SecurityCamera.Resolution,
		SecurityCamera.FieldOfView,
		SecurityCamera.RecordingDuration,
		SecurityCamera.PowerSource,
		SecurityCamera.ID,
		SecurityCamera.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.DB.QueryRowContext(ctx, query, args...).Scan(&SecurityCamera.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (s SecurityCameraModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	DELETE FROM security_cameras WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := s.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (s SecurityCameraModel) GetAll(manufacturer string, resolution string, filters Filters) ([]*SecurityCamera, error) {

	query := `
		SELECT id, created_at, manufacturer, storage_capacity, location, resolution, field_of_view, recording_duration, power_source, version 
		FROM security_cameras
		WHERE (LOWER(manufacturer) = LOWER($1) OR $1 = '')
		AND (LOWER(resolution) = LOWER($2) OR $2 = '')
		ORDER BY id
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.DB.QueryContext(ctx, query, manufacturer, resolution)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	security_cameras := []*SecurityCamera{}

	for rows.Next() {
		var securitycamera SecurityCamera
		err := rows.Scan(
			&securitycamera.ID,
			&securitycamera.CreatedAt,
			&securitycamera.Manufacturer,
			&securitycamera.StorageCapacity,
			&securitycamera.Location,
			&securitycamera.Resolution,
			&securitycamera.FieldOfView,
			&securitycamera.RecordingDuration,
			&securitycamera.PowerSource,
			&securitycamera.Version,
		)
		if err != nil {
			return nil, err
		}
		security_cameras = append(security_cameras, &securitycamera)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return security_cameras, nil

}
