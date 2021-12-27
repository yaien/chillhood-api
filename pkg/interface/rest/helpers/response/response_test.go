package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	data := map[string]string{"message": "Hello World"}
	status := http.StatusOK
	JSON(rr, data, status)
	jsondata, _ := json.Marshal(data)
	jsonstr := string(jsondata)
	if ok, _ := regexp.MatchString(jsonstr, rr.Body.String()); !ok {
		t.Errorf("expected body to be '%s', received: '%s'", rr.Body.String(), jsonstr)
	}
	if rr.Code != status {
		t.Errorf("expect response status to be %d, received: %d", status, rr.Code)
	}
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expect response content type to be application/json, received '%s'", contentType)
	}
}

func TestSend(t *testing.T) {
	rr := httptest.NewRecorder()
	data := map[string]string{"message": "Hello World"}
	status := http.StatusOK
	Send(rr, data)
	jsondata, _ := json.Marshal(data)
	jsonstr := string(jsondata)
	if ok, _ := regexp.MatchString(jsonstr, rr.Body.String()); !ok {
		t.Errorf("expected body to be '%s', received: '%s'", rr.Body.String(), jsonstr)
	}
	if rr.Code != status {
		t.Errorf("expect response status to be %d, received: %d", status, rr.Code)
	}
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expect response content type to be application/json, received '%s'", contentType)
	}
}

func TestError(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("ERROR")
	status := http.StatusBadRequest
	Error(rr, err, status)
	jsondata, _ := json.Marshal(map[string]string{"error": err.Error()})
	jsonstr := string(jsondata)
	if ok, _ := regexp.MatchString(jsonstr, rr.Body.String()); !ok {
		t.Errorf("expected body to be '%s', received: '%s'", rr.Body.String(), jsonstr)
	}
	if rr.Code != status {
		t.Errorf("expect response status to be %d, received: %d", status, rr.Code)
	}
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expect response content type to be application/json, received '%s'", contentType)
	}
}
