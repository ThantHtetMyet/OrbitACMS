package model

import "time"

type UserImage struct {
	ID              string
	UserID          string
	ImageTypeID     string
	ImageName       string
	StoredDirectory string
	IsDeleted       bool
	UploadedDate    time.Time
	UploadedBy      *string
}
