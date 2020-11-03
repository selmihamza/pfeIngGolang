package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequestByID(t *testing.T) {

	request, err := http.NewRequest("GET", "/request", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("DriverID", "11469")
	request.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRequestByID)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("404 - data not found")
	}
}
