package Models

import "time"

// FileCar struct
type FileCar struct {
	ID                 uint   `gorm:"primary_key"`
	Ref                string `gorm:"type:varchar(255)"`
	FileType           uint   `gorm:"column:file_type"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DriverCarID        uint `gorm:"column:driver_car_id"`
	FileExpirationDate time.Time
	IsValid            bool
	Name               string
	Size               int64
}
