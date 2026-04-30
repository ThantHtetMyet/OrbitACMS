package model

import "time"

type UserPermission struct {
	ID          string
	UserRoleID  string
	ModuleName  string
	CanCreate   bool
	CanRead     bool
	CanUpdate   bool
	CanDelete   bool
	IsDeleted   bool
	CreatedDate *time.Time
	UpdatedDate *time.Time
	CreatedBy   *string
	UpdatedBy   *string
}
