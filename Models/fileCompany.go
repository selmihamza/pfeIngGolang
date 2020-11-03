package Models

import "time"

// FileCompanie struct
type FileCompanie struct {
	ID       uint   `gorm:"primary_key"`
	Ref      string `gorm:"type:varchar(255)"`
	FileType uint   `gorm:"column:file_type"`
	// AdministrativeFileType AdministrativeFileType `gorm:"foreignkey:FileType;association_foreignkey:ID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DriverCompanyID uint `gorm:"column:driver_company_id"`
	// Companie            DriverCompanie `gorm:"column:driver_company_id;foreignkey:ID;association_foreignkey:driver_company_id"``gorm:"column:driver_company_id;foreignkey:ID;"`
}
