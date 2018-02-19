package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/_ah/health", nil)
	res := httptest.NewRecorder()

	healthCheckHandler(res, req)
	checkResponseCode(t, http.StatusOK, res.Code)

	if res.Body.String() != "ok" {
		t.Error("Health failed to return expected 'ok'")
	}
}

func TestPrimeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	getPrimeEndpoint(res, req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var o = &ServiceResponse{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), o); err != nil {
		t.Errorf("Error while unmarshaling object: %v", err)
		return
	}

	if o.Prime.Max != defaultMaxNumber {
		t.Errorf("Invalid response state:%v, expected: %d", o.Prime.Max, defaultMaxNumber)
		return
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
