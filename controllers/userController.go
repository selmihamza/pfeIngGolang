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
	"golang.org/x/crypto/bcrypt"
)

// Data struct
type Data struct {
	FirstName     string
	LastName      string
	Email         string
	Password      string
	Type          string
	Number        string
	HomeAddress   string
	HomeLatitude  float64
	HomeLongitude float64
	Country       string
	LicenceNumber string
	Office        uint
	SaasCompanyID uint
}

// PasswordData struct
type PasswordData struct {
	Password    string
	NewPassword string
}

// DriverInformations struct
type DriverInformations struct {
	NumberOfRideTotal          int
	NumberOfRideFinished       int
	User                       Models.BackEndUser
	SaasOfficeToDriverRelation Models.SaasOfficeToDriverRelation
	Note                       Models.DriverNote
	VehicleTypesPrimary        []Models.VehicleType
	VehicleTypesSecondary      []Models.VehicleType
}

// ProfileData struct
type ProfileData struct {
	User     Models.BackEndUser
	Driver   Models.Driver
	Addresse Models.Addresse
}

// HashPassword function
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// UpdatePassword godoc
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	driverID := params["driverID"]
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var passwordData PasswordData
	err = json.Unmarshal(body, &passwordData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	db := Helpers.DbConnect()
	var user Models.BackEndUser
	db.Model(&user).Where("driver_id = ?", driverID).First(&user)

	encryptedPassword, _ := HashPassword(passwordData.NewPassword)
	json.NewEncoder(w).Encode(encryptedPassword)
	json.NewEncoder(w).Encode(user.EncryptedPassword)

}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags back_end_users
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.BackEndUser
// @Router /register [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	db := Helpers.DbConnect()
	var count int
	db.Table("back_end_users").Where("email = ?", data.Email).Count(&count)
	if count == 0 {
		// add information in addresses table
		address := Models.Addresse{AddressDescription: data.HomeAddress, AddressZipCode: "75075", AddressLat: data.HomeLatitude, AddressLong: data.HomeLongitude, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(address)
		db.Create(&address)
		// add information in drivers table
		// DriverType == 1 ==> Manager
		driver := Models.Driver{DriverType: 1, DriverLicense: data.LicenceNumber, CreatedAt: time.Now(), UpdatedAt: time.Now(), Addresse: address, AddresseID: address.ID, DriverStatus: 0, Active: false, SignUpStep: 1}
		db.NewRecord(driver)
		db.Create(&driver)
		// create saasCompany relation with driver
		saasCompanyRelation := Models.SaasCompanyToDriverRelation{SaasCompanyID: data.SaasCompanyID, DriverID: driver.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(saasCompanyRelation)
		db.Create(&saasCompanyRelation)
		fmt.Printf("\n company id : %d", saasCompanyRelation.ID)
		// create saasOffice relation with driver
		saasOfficeRelation := Models.SaasOfficeToDriverRelation{SaasOfficeID: data.Office, DriverID: driver.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		db.NewRecord(saasOfficeRelation)
		db.Create(&saasOfficeRelation)
		fmt.Printf("\n office id : %d", saasOfficeRelation.ID)
		// add information in back_end_user table
		encryptedPassword, _ := HashPassword(data.Password)
		user := Models.BackEndUser{FirstName: data.FirstName, LastName: data.LastName, Email: data.Email, PhoneNumber: data.Number, EncryptedPassword: encryptedPassword, CreatedAt: time.Now(), UpdatedAt: time.Now(), Driver: driver, DriverID: driver.ID}
		db.NewRecord(user)
		db.Create(&user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("address mail already used!"))
	}

}

// GetDriverInformationsByID handler
// @Summary retrieve a driver information  by id
// @Description retrieve a driver information  in database  by idin request parameter
// @Tags back_end_users
// @Accept  json
// @Produce  json
// @Param id path string true "enter driver id:"
// @Success 200 {object} Models.User
// @Router /driver/{id} [get]
func GetDriverInformationsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverID := params["id"]
	db := Helpers.DbConnect()
	var user Models.BackEndUser
	db.Model(&user).Where("driver_id = ?", driverID).Find(&user)
	var saasOfficeToDriverRelation Models.SaasOfficeToDriverRelation
	db.Model(&saasOfficeToDriverRelation).Where("driver_id = ?", driverID).Find(&saasOfficeToDriverRelation)
	var note Models.DriverNote
	db.Model(&note).Where("id = ?", driverID).Find(&note)
	state := "finished"
	var count int
	var total int
	db.Table("requests").Where("driver_id = ?", driverID).Where("company_id > ?", 0).Where("state = ?", state).Count(&count)
	db.Table("requests").Where("driver_id = ?", driverID).Where("company_id > ?", 0).Count(&total)
	var car Models.DriverCar
	db.Table("driver_cars").Where("driver_id = ?", driverID).Find(&car)

	var driverCarToSasOfficeRelations Models.DriverCarToSaasOfficeRelation
	db.Table("driver_car_to_saas_office_relations").Where("driver_car_id = ?", car.ID).Find(&driverCarToSasOfficeRelations)
	var driverCarToVehicleTypeRelations []Models.DriverCarToVehicleTypeRelation
	db.Table("driver_car_to_vehicle_type_relations").Where("driver_car_to_saas_office_relation_id = ?", driverCarToSasOfficeRelations.ID).Find(&driverCarToVehicleTypeRelations)

	var ListPrimary []Models.VehicleType
	var ListSecondary []Models.VehicleType
	for i := 0; i < len(driverCarToVehicleTypeRelations); i++ {
		if driverCarToVehicleTypeRelations[i].Primary == true {
			var VehicleType Models.VehicleType
			db.Table("vehicle_types").Where("id = ?", driverCarToVehicleTypeRelations[i].VehicleTypeID).Find(&VehicleType)
			ListPrimary = append(ListPrimary, VehicleType)

		} else if driverCarToVehicleTypeRelations[i].Primary == false {
			var VehicleTypeSec Models.VehicleType
			db.Table("vehicle_types").Where("id = ?", driverCarToVehicleTypeRelations[i].VehicleTypeID).Find(&VehicleTypeSec)
			ListSecondary = append(ListSecondary, VehicleTypeSec)
		}
	}

	informations := DriverInformations{total, count, user, saasOfficeToDriverRelation, note, ListPrimary, ListSecondary}

	json.NewEncoder(w).Encode(informations)

}

// UpdateDriverStatus handler
func UpdateDriverStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverID := params["driverID"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var status int
	err = json.Unmarshal(body, &status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	db := Helpers.DbConnect()

	var saasOffice Models.SaasOfficeToDriverRelation
	db.Model(&saasOffice).Where("driver_id = ?", driverID).First(&saasOffice)
	saasOffice.DriverStatus = status
	db.Save(&saasOffice)
	var saasOfficed Models.SaasOfficeToDriverRelation
	db.Model(&saasOfficed).Where("driver_id = ?", driverID).First(&saasOfficed)
	json.NewEncoder(w).Encode(saasOfficed)

}

// GetDriverProfileByID handler
func GetDriverProfileByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverID := params["id"]
	db := Helpers.DbConnect()
	var driver Models.Driver
	db.First(&driver, driverID)
	var addresse Models.Addresse
	db.First(&addresse, driver.AddresseID)
	var user Models.BackEndUser
	db.Model(&user).Where("driver_id = ?", driver.ID).Find(&user)
	profileData := ProfileData{user, driver, addresse}
	json.NewEncoder(w).Encode(profileData)

}

// GetOfficesBySaasCompany handler
// @Summary retrieve a list of offices by saas company name
// @Description retrieve a list of offices  in database  by saas company name in request parameter
// @Tags saas_companies saas_offices
// @Accept  json
// @Produce  json
// @Param saas_company path string true "enter saas_company:"
// @Success 200 {object} Models.SaasOffice
// @Router /{saas_company}/offices [get]
func GetOfficesBySaasCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	saasParam := params["saas_company"]
	db := Helpers.DbConnect()
	var saasCompany Models.SaasCompany
	db.Model(&saasCompany).Where("name = ?", saasParam).Find(&saasCompany)
	var saasOffice []Models.SaasOffice
	db.Model(&saasOffice).Where("saas_company_id = ?", saasCompany.ID).Find(&saasOffice)

	json.NewEncoder(w).Encode(saasOffice)

}

// CompanyAndOffice struct
type CompanyAndOffice struct {
	SaasCompanyID uint
	SaasOfficeID  uint
}

// GetSaasCompanyAndOffice handler
// @Summary retrieve a saascompany and saasOffice by driverID
// @Description retrieve a saascompany and saasOffice by driverID in database with the ID in request parameter
// @Tags saas_company_to_driver_relations saas_office_to_driver_relation
// @Accept  json
// @Produce  json
// @Param driver_id path integer true "enter driver_id:"
// @Success 200 {object} CompanyAndOffice
// @Router /saascompanyAndoffice/{driver_id} [get]
func GetSaasCompanyAndOffice(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	driverID := params["driver_id"]
	var saasCompany Models.SaasCompanyToDriverRelation
	db.Model(&saasCompany).Where("driver_id = ?", driverID).First(&saasCompany)

	var saasOffice Models.SaasOfficeToDriverRelation
	db.Model(&saasOffice).Where("driver_id = ?", driverID).First(&saasOffice)
	if saasCompany.ID != 0 && saasOffice.ID != 0 {
		companyAndOffice := CompanyAndOffice{SaasCompanyID: saasCompany.SaasCompanyID, SaasOfficeID: saasOffice.SaasOfficeID}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(companyAndOffice)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}

}
