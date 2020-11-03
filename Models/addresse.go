package Models

import "time"

// Addresse struct
type Addresse struct {
	ID                 uint   `gorm:"primary_key"`
	AddressDescription string `gorm:"type:varchar(255)"`
	AddressZipCode     string `gorm:"type:varchar(255)"`
	AddressLat         float64
	AddressLong        float64
	Alias              string
	AddressComplement  string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
