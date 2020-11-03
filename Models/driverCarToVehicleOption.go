package Models

import "time"

// DriverCarToVehicleOption struct
type DriverCarToVehicleOption struct {
	ID              uint `gorm:"primary_key"`
	VehicleOptionID uint
	DriverCarID     uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Validated       uint
}
