package controllers

import (
	"driver/Helpers"
	"driver/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// PeriodData struct
type PeriodData struct {
	EventBegin string
	EventEnd   string
	DriverID   int
}

// ListAgenda struct
type ListAgenda struct {
	PrimaryTypeName   []string
	SecondaryTypeName []string
	Agenda            Models.AgendaEvent
}

// CreateAgenda function
func CreateAgenda(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var periodData PeriodData
	err = json.Unmarshal(body, &periodData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dateStart, _ := time.Parse(time.RFC3339, periodData.EventBegin)
	hrStart, minStart, secStart := dateStart.Clock()
	eventBegin := time.Date(dateStart.Year(), dateStart.Month(), dateStart.Day(), hrStart+1, minStart, secStart, dateStart.Nanosecond(), time.UTC)

	dateEnd, _ := time.Parse(time.RFC3339, periodData.EventEnd)
	hrEnd, minEnd, secEnd := dateEnd.Clock()
	eventEnd := time.Date(dateEnd.Year(), dateEnd.Month(), dateEnd.Day(), hrEnd+1, minEnd, secEnd, dateEnd.Nanosecond(), time.UTC)

	var car Models.DriverCar
	db.Model(&car).Where("driver_id = ?", periodData.DriverID).Find(&car)
	var dCtSo Models.DriverCarToSaasOfficeRelation
	db.Model(&dCtSo).Where("driver_car_id = ?", car.ID).Find(&dCtSo)

	var dCtVt []Models.DriverCarToVehicleTypeRelation
	if dCtSo.Primary == true {
		db.Model(&dCtVt).Where("driver_car_to_saas_office_relation_id = ?", dCtSo.ID).Find(&dCtVt)
	}
	var primary pq.Int64Array

	var secondary pq.Int64Array
	if dCtVt != nil {
		for i := 0; i < len(dCtVt); i++ {
			if dCtVt[i].Primary == true {
				primary = append(primary, int64(dCtVt[i].VehicleTypeID))
			} else if dCtVt[i].Primary == false {
				secondary = append(secondary, int64(dCtVt[i].VehicleTypeID))
			}
		}

	}

	agenda := Models.AgendaEvent{EventBegin: eventBegin, EventEnd: eventEnd, DriverID: periodData.DriverID, Priority: 999, PrimaryVehicleTypeIds: primary, SecondaryVehicleTypeIds: secondary, CreatedAt: time.Now(), UpdatedAt: time.Now(), Status: 0}
	db.NewRecord(agenda)
	db.Create(&agenda)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agenda)
}

// GetAgendaByID function
func GetAgendaByID(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	DriverID := params["DriverID"]

	var agenda []Models.AgendaEvent
	var listAgenda []ListAgenda
	db.Model(&agenda).Where("driver_id = ?", DriverID).Where("status != ?", 3).Limit(5).Find(&agenda)

	for i := 0; i < len(agenda); i++ {
		var typeName []string
		for j := 0; j < len(agenda[i].PrimaryVehicleTypeIds); j++ {
			var vehicleType Models.VehicleType
			db.First(&vehicleType, agenda[i].PrimaryVehicleTypeIds[j])
			typeName = append(typeName, vehicleType.Designation)
		}
		var typeNameSecondary []string
		for j := 0; j < len(agenda[i].SecondaryVehicleTypeIds); j++ {
			var vehicleType Models.VehicleType
			db.First(&vehicleType, agenda[i].SecondaryVehicleTypeIds[j])
			typeNameSecondary = append(typeNameSecondary, vehicleType.Designation)
		}
		rowAgenda := ListAgenda{typeName, typeNameSecondary, agenda[i]}
		listAgenda = append(listAgenda, rowAgenda)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(listAgenda)
}

// ConfirmPeriod function
func ConfirmPeriod(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	ID := params["ID"]

	var agenda Models.AgendaEvent

	db.Model(&agenda).Where("id = ?", ID).First(&agenda)
	agenda.Priority = 1
	db.Save(&agenda)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agenda)
}

// DeletePeriod function
func DeletePeriod(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	ID := params["ID"]

	var agenda Models.AgendaEvent

	db.Model(&agenda).Where("id = ?", ID).First(&agenda)
	agenda.Status = 3
	db.Save(&agenda)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agenda)
}

// UpdatePeriod function
func UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	ID := params["ID"]

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var periodData PeriodData
	err = json.Unmarshal(body, &periodData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dateStart, _ := time.Parse(time.RFC3339, periodData.EventBegin)
	hrStart, minStart, secStart := dateStart.Clock()
	eventBegin := time.Date(dateStart.Year(), dateStart.Month(), dateStart.Day(), hrStart+1, minStart, secStart, dateStart.Nanosecond(), time.UTC)

	dateEnd, _ := time.Parse(time.RFC3339, periodData.EventEnd)
	hrEnd, minEnd, secEnd := dateEnd.Clock()
	eventEnd := time.Date(dateEnd.Year(), dateEnd.Month(), dateEnd.Day(), hrEnd+1, minEnd, secEnd, dateEnd.Nanosecond(), time.UTC)

	var agenda Models.AgendaEvent
	db.Model(&agenda).Where("id = ?", ID).First(&agenda)
	agenda.EventBegin = eventBegin
	agenda.EventEnd = eventEnd
	agenda.UpdatedAt = time.Now()

	db.Save(&agenda)
	fmt.Print(agenda.ID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agenda)
}
