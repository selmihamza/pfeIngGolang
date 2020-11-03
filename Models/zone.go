package Models

import "time"

// Zone struct
type Zone struct {
	ID            uint `gorm:"primary_key"`
	LatMin        float64
	LatMax        float64
	LongMin       float64
	LongMax       float64
	Name          string `gorm:"type:varchar(255)"`
	Address       string `gorm:"type:varchar(255)"`
	SaasCompanyID int
	SaasOfficeID  int
	Deleted       bool
	ServiceArea   bool
	GreenZone     bool
	Type          string
	DriverZone    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
