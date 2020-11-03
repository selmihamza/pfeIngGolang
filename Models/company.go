package Models

import "time"

// DriverCompanie struct
type DriverCompanie struct {
	ID            uint   `gorm:"primary_key"`
	CompanyName   string `gorm:"type:varchar(255)"`
	CompanyDre    string `gorm:"type:varchar(255)"`
	CompanyTva    string `gorm:"type:varchar(255)"`
	CompanySiren  string `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	VatPercentage float64
	Files         []FileCompanie `gorm:"foreignkey:DriverCompanyID"`
}
