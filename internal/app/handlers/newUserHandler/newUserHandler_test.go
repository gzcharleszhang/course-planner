// +build all integration

package newUserHandler

import (
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newUserService"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx, err := testUtils.Init()
	if err != nil {
		t.Errorf("Failed to initialize test: %v\n", err)
	}
	req := utils.M{
		"first_name": "Steven",
		"last_name":  "Xu",
		"password":   "course_planner>inflight",
		"email":      "hello@stevenxu.me",
	}
	jsonStr := utils.ToRawJson(req)
	rr, err := testUtils.NewRequest("POST", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var res newUserService.Response
	decoder := json.NewDecoder(rr.Body)
	decoder.Decode(&res)
	userId := res.UserId

	// check if we can find the new user in the database
	sess, err := db.NewSession(ctx)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	defer sess.Close(ctx)
	var userData users.UserData
	err = sess.Users().FindOne(ctx, bson.M{"_id": userId}).Decode(&userData)
	if err != nil {
		t.Errorf("Cannot find the newly created user: %v", err)
	}

	// check important fields
	if !utils.StrCmp(string(userData.FirstName), req["first_name"]) ||
		!utils.StrCmp(string(userData.LastName), req["last_name"]) ||
		!utils.StrCmp(string(userData.Email), req["email"]) {
		t.Errorf("Expected %v to contain %v", utils.ToJson(userData), utils.ToJson(req))
	}

	// Test creating user with duplicate emails
	req = utils.M{
		"first_name": "Jenny",
		"last_name":  "Xu",
		"password":   "donthatemesteven",
		"email":      "hello@stevenxu.me",
	}
	jsonStr = utils.ToRawJson(req)
	rr, err = testUtils.NewRequest("POST", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Errorf("%v", err)
	}
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
