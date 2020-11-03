package Models

import "time"

// VehicleCategorie struct
type VehicleCategorie struct {
	ID            uint   `gorm:"primary_key"`
	Name          string `gorm:"type:varchar(255)"`
	TotalCapacity int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	SaasOfficeID  int
	DeletedAt     time.Time
}
