package Models

import "time"

// Customer struct
type Customer struct {
	ID                      uint `gorm:"primary_key"`
	SystemUserID            int
	CreatedAt               time.Time
	UpdatedAt               time.Time
	DefaultPassengerID      int
	FrontEndUserID          int
	SaasCompanyID           int
	CustomerCompanyName     string `gorm:"type:varchar(255)"`
	ReferralCode            string `gorm:"column:referral_code"`
	SponsorID               int
	CostCenter              string `gorm:"type:varchar(255)"`
	DefaultComment          string `gorm:"type:varchar(500)"`
	SendPdfEmail            bool
	DefaultDeliverySenderID int
	SponsorType             string
	CompanyPreferencesID    int
}
