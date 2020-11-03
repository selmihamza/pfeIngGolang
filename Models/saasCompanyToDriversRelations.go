package Models

import "time"

// SaasCompanyToDriverRelation struct
type SaasCompanyToDriverRelation struct {
	ID              uint   `gorm:"primary_key"`
	SaasCompanyID 	uint 
	DriverID     	uint 
	PhoneNumber 	string `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	EnableZendrive  bool
}
