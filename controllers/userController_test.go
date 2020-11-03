package controllers

import (
	"bytes"
	"driver/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCreateUser(t *testing.T) {

	request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(mock.UserDataCreate))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("401 - Something bad happened!  address mail must be unique")
	}
}

func TestGetOfficesBySaasCompany(t *testing.T) {
	request, err := http.NewRequest("GET", "/{saas_company}/offices", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("saas_company", "Marcel")
	request.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOfficesBySaasCompany)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
func TestGetSaasCompanyAndOffice(t *testing.T) {
	request, err := http.NewRequest("GET", "/saascompanyAndoffice/{driver_id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q, _ := url.ParseQuery(request.URL.RawQuery)
	q.Add("driver_id", "11469")
	request.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSaasCompanyAndOffice)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
