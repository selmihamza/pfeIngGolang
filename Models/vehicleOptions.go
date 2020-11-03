package Models

import "time"

// VehicleOption struct
type VehicleOption struct {
	ID                               uint   `gorm:"primary_key"`
	Name                             string `gorm:"type:varchar(255)"`
	VehicleOptionCategoryID          int
	IsDfaultForVehicleOptionCategory bool
	CreatedAt                        time.Time
	UpdatedAt                        time.Time
	SaasOfficeID                     int
	Deleted                          bool
	PriceImpact                      int
	TimeImpact                       int
	DriverOption                     bool
	AlwaysAvailable                  bool
}
