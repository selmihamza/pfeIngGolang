package Models

import (
	"time"
)

// Request struct
type Request struct {
	ID                             uint `gorm:"primary_key"`
	DriverID                       int
	InvoiceID                      int
	CustomerID                     int
	PassengerID                    int
	CompanyID                      int
	AddressPickUp                  string `gorm:"type:varchar(255)"`
	AddressPickUpLat               float64
	AddressPickUpLong              float64
	AddressDropOff                 string `gorm:"type:varchar(255)"`
	AddressDropOffLat              float64
	AddressDropOffLong             float64
	RideDate                       time.Time
	FlightNumber                   string `gorm:"type:varchar(255)"`
	Comment                        string `gorm:"type:varchar(255)"`
	PaymentType                    string `gorm:"type:varchar(255)"`
	RideType                       int
	RidePrice                      float64
	RideDiscountPercentage         float64
	EstimatePickUpDate             time.Time
	EstimateDropOffDate            time.Time
	EstimateDistanceM              int
	State                          string `gorm:"type:varchar(255)"`
	PayinID                        int
	DriverPaymentID                int
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
	SidekiqJid                     string `gorm:"type:varchar(255)"`
	RequestDiscountCode            string `gorm:"type:varchar(255)"`
	ChannelID                      int
	FlightID                       int
	DispatchDate                   time.Time
	DriverPrice                    int
	OptionsPrice                   int
	CommercialFees                 int
	DiscountCodeDiscount           int
	VehicleTypeID                  int
	DriverDiscountCodeDiscount     int
	SaasCompanyID                  int
	SaasOfficeID                   int
	CompanyPreferencesID           int
	InitialVehicleTypeID           int
	PartnerID                      string `gorm:"type:varchar(255)"`
	Closed                         bool
	OutsourcedDestinationRequestID int
	OutsourcedOriginRequestID      int
	CostCenter                     string `gorm:"type:varchar(255)"`
	ZoneID                         int
	AddressDropOffPostalCode       string `gorm:"type:varchar(255)"`
	AddressPickUpPostalCode        string `gorm:"type:varchar(255)"`
	AgendaEventID                  int
	ActiveForDispatch              bool
	PreviousRequestID              int
	RecurrentRequestID             int
	OptionIDs                      []uint8
	PriorityRequest                bool
	FireTime                       time.Time
	PriceChanged                   bool
	RideDropOffDate                time.Time
	PickUpTimeWindowID             int
	DropOffTimeWindowID            int
	Type                           string `gorm:"type:varchar(255)"`
	OnlyDriverPriceCents           int
	RuleValueTimeWindowID          int
	// CompanyReferences
	DeliveryRecipientID      int
	DeliverySenderID         int
	AddressPickUpComplement  string `gorm:"type:varchar"`
	AddressDropOffComplement string `gorm:"type:varchar"`
	DriverCarID              int
	DirectRequest            bool
	ReservationCode          string `gorm:"type:varchar"`
	// PriceComponents          JSONB
	DriverZoneID        int
	DiscountCodeID      int
	TheoreticalDuration int
	MaximumDetour       int
}
