package newTimelineHandler

import (
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newTimelineService"
	"net/http"
)

const RouteURL string = "/"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, err := users.GetUserIdFromContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var req newTimelineService.Request
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	req.UserId = userId
	res, err := newTimelineService.Execute(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}
