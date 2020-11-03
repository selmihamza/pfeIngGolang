package Models

import (
	"time"
)

// AmountsType struct
type AmountsType struct {
	BasePrice         int
	DelayFees         int
	PaidAmount        int
	DriverPrice       int
	OptionPrice       int
	CreditAmount      int
	CustomerPrice     int
	CommercialFees    int
	DiscountCodePrice int
	InvoiceTotalPrice int
}

// RequestStaticInformation struct
type RequestStaticInformation struct {
	ID                         uint `gorm:"primary_key"`
	StartDate                  time.Time
	Request                    int
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	DurationAfterFlightArrival int
	CancelDate                 time.Time
	AdminCancelDate            time.Time
	SmsMessage                 string `gorm:"type:varchar"`
	ClusterPrice               int
	WayToDate                  time.Time
	WaitDate                   time.Time
	FinishDate                 time.Time
	AcceptDate                 time.Time
	TreatAcceptedDate          time.Time
	Amounts                    AmountsType
	AdminCancelReason          string `gorm:"type:varchar"`
}
