package dto

import "time"

type UserDTO struct {
	ID          string     `json:"id"`
	UserRoleID  string     `json:"userRoleId"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	MobileNo    string     `json:"mobileNo"`
	Remark      *string    `json:"remark,omitempty"`
	LastLogin   *time.Time `json:"lastLogin,omitempty"`
	IsActive    bool       `json:"isActive"`
	IsDeleted   bool       `json:"isDeleted"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	CreatedBy   *string    `json:"createdBy,omitempty"`
	UpdatedBy   *string    `json:"updatedBy,omitempty"`
}
