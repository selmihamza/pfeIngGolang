package Models

import "time"

// DriverCarToSaasOfficeRelation struct
type DriverCarToSaasOfficeRelation struct {
	ID                uint `gorm:"primary_key"`
	DriverCarID       uint
	SaasOfficeID      int
	VehicleTypeID     int
	Validated         int
	RefusalReason     string `gorm:"type:varchar(255)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Primary           bool
	VehicleCategoryID int
}
