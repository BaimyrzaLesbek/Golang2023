package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	SecurityCameras SecurityCameraModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		SecurityCameras: SecurityCameraModel{DB: db},
	}
}
