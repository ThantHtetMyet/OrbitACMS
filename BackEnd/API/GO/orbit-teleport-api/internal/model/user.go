package model

import "time"

type User struct {
	ID            string
	UserRoleID    string
	FirstName     string
	LastName      string
	Email         string
	MobileNo      string
	LoginPassword string
	Remark        *string
	LastLogin     *time.Time
	IsActive      bool
	IsDeleted     bool
	CreatedDate   time.Time
	UpdatedDate   time.Time
	CreatedBy     *string
	UpdatedBy     *string
}
