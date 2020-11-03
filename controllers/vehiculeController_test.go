package controllers

import (
	"bytes"
	"driver/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateVehicle(t *testing.T) {

	request, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(mock.VehicleDataCreate))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "multipart/form-data; boundary=<calculated when request")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateVehicle)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
func TestGetVehicleTypes(t *testing.T) {

	request, err := http.NewRequest("GET", "/vehicule_types", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("saas_officeID", "1")
	request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVehicleTypes)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("404 - data not found")
	}
}
func TestGetVehicleOptions(t *testing.T) {

	request, err := http.NewRequest("GET", "/vehicule_options", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("saas_officeID", "1")
	request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVehicleOptions)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("404 - data not found")
	}
}
func TestGetVehicleCategories(t *testing.T) {

	request, err := http.NewRequest("GET", "/vehicule_categories", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("saas_officeID", "1")
	request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVehicleCategories)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("404 - data not found")
	}
}
func TestGetVehicleZones(t *testing.T) {

	request, err := http.NewRequest("GET", "/vehicule_zones/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	// q := request.URL.Query()
	// q.Add("saas_officeID", "1")
	// request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVehicleZones)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("404 - data not found")
	}
}
