package Models

import "time"

// SaasOffice struct
type SaasOffice struct {
	ID            uint   `gorm:"primary_key"`
	Name          string `gorm:"type:varchar(255)"`
	SaasCompanyID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
