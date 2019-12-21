package getTimelineHandler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getTimelineService"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newTimelineService"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx, user, err := testUtils.InitWithUser()
	if err != nil {
		t.Errorf("Failed to initialize test: %v\n", err)
	}
	defer testUtils.CleanUp()

	ctx = context.WithValue(ctx, contextKeys.UserIdKey, user.Id)

	// set up test Timeline
	req := newTimelineService.Request{
		UserId: user.Id,
		Name:   timelines.TimelineName("test_timeline_1"),
	}
	tlRes, err := newTimelineService.Execute(ctx, req)
	if err != nil {
		t.Error(err)
	}

	timeline := tlRes.Timeline

	// get timeline by id
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("timeline_id", string(timeline.Id))
	ctx = context.WithValue(ctx, chi.RouteCtxKey, rctx)
	jsonStr := utils.ToRawJson(utils.M{})
	rr, err := testUtils.NewRequest(ctx, "GET", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\nerror: %v",
			status, http.StatusOK, testUtils.GetErrorResponse(rr))
	}
	var res getTimelineService.Response
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&res)
	if err != nil {
		t.Error(err)
	}
	tl := res.Timeline
	if tl.Id != timeline.Id || tl.Name != timeline.Name {
		t.Errorf("expected response to be %v, got %v instead",
			utils.ToJson(timeline), utils.ToJson(tl))
	}

	// get timeline by fake id
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("timeline_id", "fake_id")
	ctx = context.WithValue(ctx, chi.RouteCtxKey, rctx)
	jsonStr = utils.ToRawJson(utils.M{})
	rr, err = testUtils.NewRequest(ctx, "GET", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	// expects mongo no document error
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
	errorStr := testUtils.GetErrorResponse(rr)
	if !strings.Contains(errorStr, mongo.ErrNoDocuments.Error()) {
		t.Errorf("expected error to be %v, got %v",
			mongo.ErrNoDocuments.Error(), errorStr)
	}
}
