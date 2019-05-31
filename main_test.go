package main

import (
  "log"
	"testing"
  "encoding/json"
  "net/http"
  "net/http/httptest"
)

func TestResponseForStations(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "" || body == "[]" {
		t.Errorf("Expected values in the array. Got nil")
	}
}

func TestResponseForStationsWithPage(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations?page=2", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
  stationResponses := make([]StationResponse,0)
  json.Unmarshal([]byte(response.Body.String()),&stationResponses)
  log.Println(len(stationResponses))
	if (len(stationResponses) > 20) {
		t.Errorf("A page cannot have more than 20 values")
	}
}

func TestResponseForStationsInService(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations/in-service", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
  stationResponses := make([]StationResponse,0)
  json.Unmarshal([]byte(response.Body.String()),&stationResponses)
	if (len(stationResponses) <= 0) {
		t.Errorf("There are stations in service")
	}
}

func TestResponseForStationsNotInService(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations/not-in-service", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
  stationResponses := make([]StationResponse,0)
  json.Unmarshal([]byte(response.Body.String()),&stationResponses)
	if (len(stationResponses)  <= 0) {
		t.Errorf("There are stations that are not in service")
	}
}

func TestResponseForPageNumberOutsideLimit(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations/not-in-service?page=34", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnprocessableEntity, response.Code)

}

func TestResponseForInvalidPage(t *testing.T) {

	req, _ := http.NewRequest("GET", "/stations/not-in-service?page=test", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestResponseForInvalidStationIdForDocakable(t *testing.T) {

	req, _ := http.NewRequest("GET", "/dockable/1987690/2", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnprocessableEntity, response.Code)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	NewRouter().ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
