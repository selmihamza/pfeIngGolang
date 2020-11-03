package Models

import "time"

// VehicleType struct
type VehicleType struct {
	ID            uint   `gorm:"primary_key"`
	Designation   string `gorm:"type:varchar(255)"`
	IsAvailable   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	KmPrice       int
	MinPrice      int
	PickupPrice   int
	TypeCode      int
	SaasCompanyID int
	SaasOfficeID int
	MaxPlace      int
	LogoVehicle   string `gorm:"type:varchar(255)"`
}
