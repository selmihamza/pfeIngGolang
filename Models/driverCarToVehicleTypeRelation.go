package Models

import "time"

// DriverCarToVehicleTypeRelation struct
type DriverCarToVehicleTypeRelation struct {
	ID                              uint `gorm:"primary_key"`
	DriverCarToSaasOfficeRelationID uint
	Validated                       int
	RefusalReason                   string `gorm:"type:varchar(255)"`
	Primary                         bool
	VehicleTypeID                   uint
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}
