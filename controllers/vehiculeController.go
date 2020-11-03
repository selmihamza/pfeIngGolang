package controllers

import (
	"driver/Helpers"
	"driver/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetVehicleTypes Handler
// @Summary retrieve a Vehicule types by saas office ID
// @Description retrieve a Vehicule types in database with the saas office ID
// @Tags vehicle_types
// @Accept  json
// @Produce  json
// @Param saas_officeID path int true "enter saas_officeID:"
// @Success 200 {object} Models.VehicleType
// @Router /vehicle_types/{saas_officeID} [get]
func GetVehicleTypes(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	saasParam := params["saas_officeID"]
	var types []Models.VehicleType
	db.Where("saas_office_id = ?", saasParam).Where("is_available = ?", true).Where("deleted = ?", false).Find(&types)
	if len(types) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}

}

// GetVehicleCategories Handler
// @Summary retrieve a Vehicule categories by saas office ID
// @Description retrieve a Vehicule categories in database with the saas office ID
// @Tags vehicle_categorie
// @Accept  json
// @Produce  json
// @Param saas_officeID path int true "enter saas_officeID:"
// @Success 200 {object} Models.VehicleCategorie
// @Router /vehicle_categories/{saas_officeID} [get]
func GetVehicleCategories(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	saasParam := params["saas_officeID"]
	var categories []Models.VehicleCategorie
	db.Where("saas_office_id = ?", saasParam).Find(&categories)
	if len(categories) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categories)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}

}

// GetVehicleOptions Handler
// @Summary retrieve a Vehicule options by saas office ID
// @Description retrieve a Vehicule options in database with the office ID
// @Tags vehicle_options
// @Accept  json
// @Produce  json
// @Param saas_officeID path int true "enter saas_officeID:"
// @Success 200 {object} Models.VehicleOption
// @Router /vehicle_options/{saas_officeID} [get]
func GetVehicleOptions(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	saasParam := params["saas_officeID"]
	var options []Models.VehicleOption
	db.Where("saas_office_id = ?", saasParam).Where("deleted = ?", false).Where("always_available = ?", false).Find(&options)
	if len(options) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(options)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}

}

// GetVehicleZones Handler
// @Summary retrieve a list of zones by saas office id
// @Description  retrieve a list of zones  in database by saas office id in request parameter
// @Tags zones
// @Accept  json
// @Produce  json
// @Param saas_officeID path string true "enter saas_office ID:"
// @Success 200 {object} Models.Zone
// @Router /vehicle_zones/{saas_officeID} [get]
func GetVehicleZones(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	saasParam := params["saas_officeID"]
	var zones []Models.Zone
	db.Where("saas_office_id = ?", saasParam).Where("driver_zone = ?", true).Where("deleted = ?", false).Find(&zones)
	if len(zones) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(zones)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}

}

// CreateVehicle Handler
// @Summary Create a new vehicle
// @Description Create a new vehicle with the input payload
// @Tags driver_cars
// @Accept  mpfd
// @Produce  json
// @Param DriverID query int true "enter driver id (tester avec l'id : 11469 ) : "
// @Param DriverCompanyID query int true "enter driver company id (tester avec l'id : 10725 ) : "
// @Param SaasOfficeID query int true "enter saas office id (tester avec l'id : 1 ) : "
// @Param Color query string false "enter Color:"
// @Param Matricule query string false "enter Matricule:"
// @Param Year query string false "enter Year:"
// @Param Model query string false "enter Model:"
// @Param Mark query string false "enter Mark:"
// @Param ASSUR formData file false "enter assur file:"
// @Param CGV formData file false "enter cgv file:"
// @Param ATR formData file false "enter atr file:"
// @Success 200 {object} Models.DriverCar
// @Router /vehicle [post]
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()

	// add information in driver_cars table
	Year, _ := strconv.Atoi(r.FormValue("Year"))
	DriverID, _ := strconv.Atoi(r.FormValue("DriverID"))
	DriverCompanyID, _ := strconv.Atoi(r.FormValue("DriverCompanyID"))
	// CarType  r.FormValue("Location") r.FormValue("CarInUse")
	vehicle := Models.DriverCar{Color: r.FormValue("Color"), Matricule: r.FormValue("Matricule"), Year: Year, CarModel: r.FormValue("Model"), CarMarque: r.FormValue("Mark"), CreatedAt: time.Now(), UpdatedAt: time.Now(), DriverCompanyID: DriverCompanyID, DriverID: DriverID, Validated: 0}
	db.NewRecord(vehicle)
	db.Create(&vehicle)
	// add information in driver_car_to_saas_office_relation table
	saasOfficeID, _ := strconv.Atoi(r.FormValue("SaasOfficeID"))
	vehicleCategoryID, _ := strconv.Atoi(r.FormValue("VehicleCategoryID"))
	driverCarToSaasOfficeRelation := Models.DriverCarToSaasOfficeRelation{DriverCarID: vehicle.ID, SaasOfficeID: saasOfficeID, Validated: 0, Primary: true, VehicleCategoryID: vehicleCategoryID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.NewRecord(driverCarToSaasOfficeRelation)
	db.Create(&driverCarToSaasOfficeRelation)
	// add information in driver_car_to_vehicle_type_relation table
	var primary []Models.VehicleType
	json.Unmarshal([]byte(r.FormValue("VehicleTypesPrimary")), &primary)
	for i := 0; i < len(primary); i++ {
		driverCarToVehicleTypeRelation := Models.DriverCarToVehicleTypeRelation{DriverCarToSaasOfficeRelationID: driverCarToSaasOfficeRelation.ID, Validated: 0, Primary: true, VehicleTypeID: primary[i].ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(driverCarToVehicleTypeRelation)
		db.Create(&driverCarToVehicleTypeRelation)
	}
	var secondary []Models.VehicleType
	json.Unmarshal([]byte(r.FormValue("VehicleTypesSecondary")), &secondary)
	for i := 0; i < len(secondary); i++ {
		driverCarToVehicleTypeRelationSecondary := Models.DriverCarToVehicleTypeRelation{DriverCarToSaasOfficeRelationID: driverCarToSaasOfficeRelation.ID, Validated: 0, Primary: false, VehicleTypeID: secondary[i].ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(driverCarToVehicleTypeRelationSecondary)
		db.Create(&driverCarToVehicleTypeRelationSecondary)
	}
	// add information in driver_car_to_vehicle_option table
	var options []Models.VehicleOption
	json.Unmarshal([]byte(r.FormValue("VehicleOptions")), &options)
	for i := 0; i < len(options); i++ {
		driverCarToVehicleOptions := Models.DriverCarToVehicleOption{VehicleOptionID: options[i].ID, DriverCarID: vehicle.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(driverCarToVehicleOptions)
		db.Create(&driverCarToVehicleOptions)
	}
	// add information in file_cars table
	//---START UPLOAD ASSUR FILE--------------------------------------------------------------
	fileCarAssur := Helpers.UploadFile(r, "assur", vehicle.ID, 0)
	db.NewRecord(fileCarAssur)
	db.Create(&fileCarAssur)
	//---START UPLOAD CGV FILE---------------------------------------------------------------
	fileCarCGV := Helpers.UploadFile(r, "cgv", vehicle.ID, 1)
	db.NewRecord(fileCarCGV)
	db.Create(&fileCarCGV)
	//---START UPLOAD ATR FILE---------------------------------------------------------------
	fileCarATR := Helpers.UploadFile(r, "atr", vehicle.ID, 2)
	db.NewRecord(fileCarATR)
	db.Create(&fileCarATR)

	w.WriteHeader(http.StatusOK)
	fmt.Printf("vehicle.ID : %d", vehicle.ID)
	json.NewEncoder(w).Encode(vehicle.ID)
}
