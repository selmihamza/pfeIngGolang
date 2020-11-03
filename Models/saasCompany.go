package Models

import "time"

// SaasCompany struct
type SaasCompany struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SaasOffices []SaasOffice
}
