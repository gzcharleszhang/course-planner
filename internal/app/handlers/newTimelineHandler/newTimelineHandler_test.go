// +build all integration

package newTimelineHandler

import (
	"context"
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils/testUtils"
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

	// Creating a timeline with no plans or courses
	timelineName := timelines.TimelineName("timeline1")
	req := utils.M{
		"name": timelineName,
	}
	jsonStr := utils.ToRawJson(req)
	rr, err := testUtils.NewRequest(ctx, "POST", RouteURL, jsonStr, Handler)
	if err != nil {
		t.Error(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\nerror: %v",
			status, http.StatusOK, testUtils.GetErrorResponse(rr))
	}
	var res newTimelineService.Response
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&res)
	if err != nil {
		t.Error(err)
	}
	if res.Timeline.Name != timelineName {
		t.Errorf("expected timeline name to be %v, got %v",
			timelineName, res.Timeline.Name)
	}
}
