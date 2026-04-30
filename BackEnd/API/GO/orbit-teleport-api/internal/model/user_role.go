package model

import "time"

type UserRole struct {
	ID          string
	RoleName    string
	Description *string
	IsDeleted   bool
	CreatedDate *time.Time
	UpdatedDate *time.Time
	CreatedBy   *string
	UpdatedBy   *string
}
