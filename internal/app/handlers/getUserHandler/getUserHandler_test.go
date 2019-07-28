// +build all integration

package getUserHandler

import (
	"context"
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getUserService"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newUserService"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx, err := testUtils.Init()
	if err != nil {
		t.Errorf("Failed to initialize test: %v\n", err)
	}
	req := newUserService.Request{
		FirstName: "Steven",
		LastName:  "Xu",
		Email:     "hello@stevenxu.me",
		Password:  "mrcalcaward",
	}
	res, err := newUserService.Execute(ctx, req)
	if err != nil {
		t.Errorf("Failed to create new user: %v\n", err)
	}
	ctx = context.WithValue(ctx, contextKeys.UserIdKey, res.UserId)
	ctx = context.WithValue(ctx, contextKeys.UserRoleKey, roles.NewConrad())
	getReq := utils.M{
		"user_id": res.UserId,
	}
	jsonStr := utils.ToRawJson(getReq)
	rr, err := testUtils.NewRequest(ctx, "GET", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var getRes getUserService.Response
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&getRes)
	if err != nil {
		t.Error(err)
	}
	userRes := getRes.User
	// check important fields
	if userRes.FirstName != req.FirstName ||
		userRes.LastName != req.LastName ||
		userRes.Email != req.Email {
		t.Errorf("Expected %v to contain %v", utils.ToJson(userRes), utils.ToJson(req))
	}

	// getting an user that doesn't exist
	getReq = utils.M{
		"user_id": "abc123",
	}
	jsonStr = utils.ToRawJson(getReq)
	rr, err = testUtils.NewRequest(ctx, "GET", RouteURL, jsonStr, Handler)
	if err == nil {
		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusInternalServerError)
		}
	} else {
		t.Error("Expected error")
	}

}
