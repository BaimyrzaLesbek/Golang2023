package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	SecurityCamera SecurityCameraModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		SecurityCamera: SecurityCameraModel{DB: db},
	}
}