package Models

import "time"

// BackEndUser struct
type BackEndUser struct {
	ID                uint   `gorm:"primary_key"`
	FirstName         string `gorm:"type:varchar(255)"`
	LastName          string `gorm:"type:varchar(255)"`
	Email             string `gorm:"type:varchar(255);unique_index"`
	PhoneNumber       string `gorm:"type:varchar(255)"`
	EncryptedPassword string `gorm:"size:255"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Locale            string `gorm:"size:255"`
	DriverID          uint   `gorm:"column:driver_id;foreignkey:ID;"`
	Driver            Driver `gorm:"column:driver_id;foreignkey:ID;association_foreignkey:driver_id"`
}

