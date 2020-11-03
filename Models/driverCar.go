package Models

import "time"

// DriverCar struct
type DriverCar struct {
	ID              uint   `gorm:"primary_key"`
	Color           string `gorm:"type:varchar(255)"`
	Matricule       string `gorm:"type:varchar(255)"`
	CarType         int
	CarStatus       int
	Year            int
	CarModel        string `gorm:"type:varchar(255)"`
	CarMarque       string `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DriverCompanyID int
	IsFr            bool
	VehicleTypeID   int
	DriverID        int
	Validated       int
	RefusalReason   string `gorm:"type:varchar(255)"`
	Deleted         bool
}
