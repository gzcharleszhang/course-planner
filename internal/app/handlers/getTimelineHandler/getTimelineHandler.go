package getTimelineHandler

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getTimelineService"
	"net/http"
)

const RouteURL string = "/{timeline_id}"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	timelineId := timelines.TimelineId(chi.URLParam(r, "timeline_id"))
	req := getTimelineService.Request{TimelineId: timelineId}
	res, err := getTimelineService.Execute(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}
