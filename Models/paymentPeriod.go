package Models

import "time"

// PaymentPeriod struct
type PaymentPeriod struct {
	ID         uint `gorm:"primary_key"`
	BeginDate  time.Time
	EndDate    time.Time
	PeriodType int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
