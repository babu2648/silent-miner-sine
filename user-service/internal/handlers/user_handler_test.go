package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"user-service/internal/models"
)

func TestGetUserByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expectedUser := models.User{
		ID:    "1",
		PayID: "user1@superapp",
		Name:  "Test User",
	}
	var actualUser models.User
	err = json.NewDecoder(rr.Body).Decode(&actualUser)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if !reflect.DeepEqual(actualUser, expectedUser) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actualUser, expectedUser)
	}
}