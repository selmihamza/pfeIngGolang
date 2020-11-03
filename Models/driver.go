package Models

import "time"

// Driver struct
type Driver struct {
	ID                    uint `gorm:"primary_key"`
	DriverType            int
	DriverLicense         string `gorm:"type:varchar(255)"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	AddresseID            uint     `gorm:"column:driver_main_address_id"`
	Addresse              Addresse `gorm:"association_foreignkey:driver_main_address_id"`
	DriverStatus          int
	Active                bool
	SignUpStep            int
	DriverCompanyID       uint           `gorm:"column:driver_company_id"`
	DriverCompanies       DriverCompanie `gorm:"column:driver_company_id;foreignkey:ID;association_foreignkey:driver_company_id"`
	CurrentDriverCarID    int
	DriverInsuranceNumber string `gorm:"column:driver_insurance_number"`
	ReferralCode          string `gorm:"column:referral_code"`
	// LastTraceDate
	// LastTraceLat
	// LastTracelong
	// LastTraceState
	// LastTraceRideId
	// LastStateDate
	// Priority
	// ValidateAccount
	// Affected
	// Completed
	// Notes
	// paymentperiodType
	// AccountState
	// AssignedState
}
