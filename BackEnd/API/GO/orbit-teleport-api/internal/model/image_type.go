package model

import "time"

type ImageType struct {
	ID            string
	ImageTypeName string
	IsDeleted     bool
	CreatedDate   time.Time
	UpdatedDate   time.Time
	CreatedBy     *string
	UpdatedBy     *string
}
