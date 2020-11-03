package Models

import "time"

// DriverInvoice struct
type DriverInvoice struct {
	ID                                  uint `gorm:"primary_key"`
	PeriodID                            int
	InvoiceNumber                       string `gorm:"type:varchar(255)"`
	PayoutTransactionID                 int
	DriverID                            int
	InitialInvoiceTotal                 int
	InitialDriverTotal                  int
	CompanyFeesPercentage               float64
	TvaPercentage                       float64
	TvaNumber                           string `gorm:"type:varchar(255)"`
	CreatedAt                           time.Time
	UpdatedAt                           time.Time
	InvoiceVersion                      int
	SaasCompanyFeesPaymentTransactionID int
	SaasOfficeID                        int
	AmountPaidByCash                    int
	PeriodBonusMalus                    int
	CommercialFees                      int
	Paid                                bool
	InvoiceCount                        int
	CompanyFeesCents                    int
	CompanyFeesVatPercentage            float64
	CompanyMarkupVatPercentage          float64
	ResponsibleID                       int
	DriverCompanyID                     int
}
