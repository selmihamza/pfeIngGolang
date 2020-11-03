package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	var jsonStr = []byte(`{"email": "selmi.hamza@gmail.com","password": "1234565"}`)
	mux := http.NewServeMux()
	mux.HandleFunc("/login", Auth)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &result)

	fmt.Println(result["success"])
	expected := true

	if result["success"] != expected {
		t.Errorf("handler returned unexpected : got %v want %v",
			result["success"], expected)
	}
}
