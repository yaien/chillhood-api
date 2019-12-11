package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type nextspy struct {
	called bool
}

func (n *nextspy) handle(w http.ResponseWriter, r *http.Request) {
	n.called = true
}

type structure struct {
	Name      string `validate:"required"`
	Existence int    `validate:"required,gt=0"`
}

func (s *structure) Validate() map[string]string {
	validator := GetValidator()
	errors, _ := validator.Validate(s)
	fmt.Println(errors)
	return errors
}

func getbody(i interface{}) (io.Reader, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func TestInvalidFor(t *testing.T) {
	handler := NewHandler(&structure{})
	recorder := httptest.NewRecorder()
	next := &nextspy{}
	body, err := getbody(&structure{"", 0})
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest("POST", "/post", body)
	handler(recorder, request, next.handle)
	result := recorder.Result()

	if result.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code to be %d, received %v", http.StatusBadRequest, result.StatusCode)
	}

	var errors map[string]string
	json.NewDecoder(result.Body).Decode(&errors)
	if errors == nil || len(errors) == 0 {
		t.Errorf("expected validation to have errors, received %v", errors)
	}

}
