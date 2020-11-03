package controllers

import (
	"driver/Helpers"
	"driver/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Path struct
type Path struct {
	Lat float64
	Lng float64
}

// ReturnPath struct
type ReturnPath struct {
	Distance          float64
	Duration          float64
	DurationInTraffic float64
	Path              []Path
}

// PathAndRequest struct
type PathAndRequest struct {
	Request                  Models.Request
	PathData                 ReturnPath
	RequestStaticInformation Models.RequestStaticInformation
}

// Rides struct
type Rides struct {
	Request      Models.Request
	FrontEndUser Models.FrontEndUser
}

// GetRequestByID function
// @Summary retrieve a request by driver ID
// @Description retrieve a request in database with the driver ID
// @Tags request
// @Accept  json
// @Produce  json
// @Param DriverID path int true "enter driver id :11469:"
// @Success 200 {object} Models.Request
// @Router /request/{DriverID} [get]
func GetRequestByID(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	DriverID := params["DriverID"]
	var request Models.Request
	db.Where("driver_id = ?", DriverID).Where("active_for_dispatch = ?", false).Order("ride_date desc").First(&request)

	url := fmt.Sprintf("https://maps-manager.herokuapp.com/api/directions?provider=google&origin=%f,%f&destination=%f,%f&traffic=true&WithPath=true", request.AddressPickUpLat, request.AddressPickUpLong, request.AddressDropOffLat, request.AddressDropOffLong)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTU1NjkwMDQsInVzZXIiOiJ5dXNvIn0.EgoSndzYbmU_vrWTjF93jISz0XyQXwehPh15mN-T9J0"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Token", token)
	resp, err := http.DefaultClient.Do(req) // Call the API
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	var pathreturn ReturnPath
	if resp.StatusCode == http.StatusOK {
		path, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error to get result from api")
		}
		//Print the Result that is returned in the response of the API
		errpath := json.Unmarshal(path, &pathreturn)
		if errpath != nil {
			fmt.Println("error path", errpath)
		}
	} else {
		fmt.Println("StatusCode is not ok")
	}
	var requestStaticInformation Models.RequestStaticInformation
	db.Model(&requestStaticInformation).Where("request_id = ?", request.ID).First(&requestStaticInformation)
	pathAndRequest := PathAndRequest{request, pathreturn, requestStaticInformation}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pathAndRequest)

}

// GetListRequestByID function
// @Summary retrieve a request by driver ID
// @Description retrieve a request in database with the driver ID
// @Tags request
// @Accept  json
// @Produce  json
// @Param DriverID path int true "enter driver id :11469:"
// @Success 200 {object} Models.Request
// @Router /request/{DriverID} [get]
func GetListRequestByID(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	DriverID := params["DriverID"]
	var requests []Models.Request
	var ListRides []Rides
	db.Where("driver_id = ?", DriverID).Where("agenda_event_id > ?", 0).Order("ride_date desc").Find(&requests)

	for i := 0; i < len(requests); i++ {
		var costumer Models.Customer
		db.First(&costumer, requests[i].CustomerID)
		var frontUser Models.FrontEndUser
		db.First(&frontUser, costumer.FrontEndUserID)
		ride := Rides{requests[i], frontUser}
		ListRides = append(ListRides, ride)
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(len(requests))
	json.NewEncoder(w).Encode(ListRides)

}
