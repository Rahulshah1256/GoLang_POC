package main

import (
	"time"
)

// Admission represents an admission form.
type Admission struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	FullName        string    `validate:"required" json:"fullName"`
	Gender          string    `validate:"required" json:"gender"`
	Dob             time.Time `validate:"required" json:"dob"`
	Address         string    `validate:"required" json:"address"`
	PhoneNo         string    `validate:"required" json:"phoneNo"`
	PreviousCompany string    `validate:"required" json:"previousCompany"`
	Designation     string    `validate:"required" json:"designation"`
	Skills          string    `validate:"required" json:"skills"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// TableName specifies the table name for the Admission model.
func (Admission) TableName() string {
	return "admissions"
}
