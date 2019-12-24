// +build all integration

package getTimelinesHandler

import (
	"context"
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getTimelinesService"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newTimelineService"
	"net/http"
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
	timelineName := timelines.TimelineName("test_timeline")
	req := newTimelineService.Request{
		UserId: user.Id,
		Name:   timelineName,
	}
	tlRes, err := newTimelineService.Execute(ctx, req)
	if err != nil {
		t.Error(err)
	}

	timeline := tlRes.Timeline

	// get all timelines for this user
	jsonStr := utils.ToRawJson(utils.M{})
	rr, err := testUtils.NewRequest(ctx, "GET", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\nerror: %v",
			status, http.StatusOK, testUtils.GetErrorResponse(rr))
	}
	var res getTimelinesService.Response
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&res)
	if err != nil {
		t.Error(err)
	}
	tls := res.Timelines
	if len(tls) != 1 || (*tls[0]).Id != timeline.Id || (*tls[0]).Name != timeline.Name {
		t.Errorf("expected response to contain %v, got %v instead",
			utils.ToJson(timeline), utils.ToJson(tls))
	}
}
