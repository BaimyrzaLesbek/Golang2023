package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	SecurityCameras SecurityCameraModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		SecurityCameras: SecurityCameraModel{DB: db},
	}
}
