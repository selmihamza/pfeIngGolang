package Models

import (
	"time"
)

// SaasOfficeToDriverRelation struct
type SaasOfficeToDriverRelation struct {
	ID                              uint `gorm:"primary_key"`
	SaasOfficeID                    uint
	DriverID                        uint
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
	DriverStatus                    int
	ValidateAccount                 int
	Priority                        int
	Completed                       string `gorm:"type:varchar(255)"`
	Notes                           string
	PaymentPeriodType               int
	AccountState                    int
	LastTraceDate                   time.Time
	LastStateDate                   time.Time
	LastTraceLat                    float64
	LastTraceLong                   float64
	LastTraceState                  int
	LastTraceRideID                 int
	Active                          bool
	DefaultCommission               float64
	NoAgenda                        bool
	ValidDocument                   bool
	DriverHumanID                   string `gorm:"type:varchar(255)"`
	CommissionPercentage            float64
	CommissionAmountPerInvoiceCents int
	CommissionAmountPerRideCents    int
	CanEditRegularAgenda            bool
	CommunicationConsentDate        string
	LastAppUpdate                   time.Time
}
