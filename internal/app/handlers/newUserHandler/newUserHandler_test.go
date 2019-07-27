// +build all integration

package newUserHandler

import (
	"bytes"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecute(t *testing.T) {
	_, err := utils.InitTest()
	if err != nil {
		t.Errorf("Failed to initialize test: %v\n", err)
	}
	jsonStr := utils.M{
		"first_name": "Steven",
		"last_name":  "Xu",
		"password":   "course_planner>inflight",
		"email":      "hello@stevenxu.me",
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(utils.ToRawJson(jsonStr)))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	fmt.Print(expected)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
