// +build all integration

package getUserHandler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
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
	// adding url param
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", string(res.UserId))
	ctx = context.WithValue(ctx, contextKeys.UserIdKey, res.UserId)
	ctx = context.WithValue(ctx, contextKeys.UserRoleKey, roles.NewConrad())
	ctx = context.WithValue(ctx, chi.RouteCtxKey, rctx)
	rr, err := testUtils.NewRequest(ctx, "GET", RouteURL, []byte{}, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var getRes utils.M
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&getRes)
	if err != nil {
		t.Error(err)
	}
	userRes := getRes["user"].(map[string]interface{})
	// check important fields
	if !utils.StrCmp(string(req.FirstName), userRes["first_name"]) ||
		!utils.StrCmp(string(req.LastName), userRes["last_name"]) ||
		!utils.StrCmp(string(req.Email), userRes["email"]) {
		t.Errorf("Expected %v to contain %v", utils.ToJson(getRes), utils.ToJson(req))
	}

	// getting an user that doesn't exist
	// adding url param
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("user_id", "abc123")
	ctx = context.WithValue(ctx, contextKeys.UserIdKey, users.UserId("abc123"))
	ctx = context.WithValue(ctx, contextKeys.UserRoleKey, roles.NewConrad())
	ctx = context.WithValue(ctx, chi.RouteCtxKey, rctx)
	rr, err = testUtils.NewRequest(ctx, "GET", RouteURL, []byte{}, Handler)
	if err == nil {
		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusInternalServerError)
		}
	} else {
		t.Error("Expected error")
	}
}
