package Models

import "time"

// AdministrativeFileType struct
type AdministrativeFileType struct {
	ID          uint   `gorm:"primary_key"`
	Designation string `gorm:"type:varchar(255)"`
	Code        string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Active      bool
}
