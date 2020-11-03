package Models

import (
	"time"

	"github.com/lib/pq"
)

// AgendaEvent struct
type AgendaEvent struct {
	ID              uint `gorm:"primary_key"`
	UserAgendasID   int
	EventType       int
	EventLabel      string `gorm:"type:varchar(255)"`
	EventParameters string `gorm:"type:varchar(255)"`
	EventBegin      time.Time
	EventEnd        time.Time
	Priority        int
	DriverID        int
	Status          int
	// DispatchDate        p
	VehicleTypeID           int
	SaasOfficeID            int
	CurrentPeriodWeight     float64
	CreatedAt               time.Time
	UpdatedAt               time.Time
	StartLocationLat        float64
	StartLocationLong       float64
	EndLocationLat          float64
	EndLocationLong         float64
	PrimaryVehicleTypeIds   pq.Int64Array
	SecondaryVehicleTypeIds pq.Int64Array
	VehicleOptionIds        pq.Int64Array
	// DriverZonesInfo jsonb
	PrimaryDispatchCategories   pq.StringArray
	SecondaryDispatchCategories pq.StringArray
	// DriverCarID

}
