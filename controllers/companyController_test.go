package controllers

import (
	"bytes"
	"driver/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCompany(t *testing.T) {

	request, err := http.NewRequest("POST", "/company", bytes.NewBuffer(mock.CompanyDataCreate))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "multipart/form-data; boundary=<calculated when request")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetCompanyByID(t *testing.T) {
	request, err := http.NewRequest("GET", "/company", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCompanyByID)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestUpdateCompany(t *testing.T) {
	request, err := http.NewRequest("PUT", "/company", bytes.NewBuffer(mock.CompanyDataUpdate))
	if err != nil {
		t.Fatal(err)
	}
	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateCompany)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
