package Models

import "time"

// DriverNote struct
type DriverNote struct {
	ID        uint `gorm:"primary_key"`
	RequestID int
	Note      int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ChanelID  int
}
