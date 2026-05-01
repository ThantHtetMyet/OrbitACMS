package dto

import "time"

type HealthResponse struct {
	Status string `json:"status"`
	API    string `json:"api"`
}

type ImageTypeDTO struct {
	ID            string    `json:"id"`
	ImageTypeName string    `json:"imageTypeName"`
	IsDeleted     bool      `json:"isDeleted"`
	CreatedDate   time.Time `json:"createdDate"`
	UpdatedDate   time.Time `json:"updatedDate"`
}

type UserImageDTO struct {
	ID              string    `json:"id"`
	UserID          string    `json:"userId"`
	ImageTypeID     string    `json:"imageTypeId"`
	ImageName       string    `json:"imageName"`
	StoredDirectory string    `json:"storedDirectory"`
	IsDeleted       bool      `json:"isDeleted"`
	UploadedDate    time.Time `json:"uploadedDate"`
}

type UserPermissionDTO struct {
	ID          string     `json:"id"`
	UserRoleID  string     `json:"userRoleId"`
	ModuleName  string     `json:"moduleName"`
	CanCreate   bool       `json:"canCreate"`
	CanRead     bool       `json:"canRead"`
	CanUpdate   bool       `json:"canUpdate"`
	CanDelete   bool       `json:"canDelete"`
	IsDeleted   bool       `json:"isDeleted"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	UpdatedDate *time.Time `json:"updatedDate,omitempty"`
}

type UserRoleDTO struct {
	ID          string     `json:"id"`
	RoleName    string     `json:"roleName"`
	IsDeleted   bool       `json:"isDeleted"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	UpdatedDate *time.Time `json:"updatedDate,omitempty"`
}
