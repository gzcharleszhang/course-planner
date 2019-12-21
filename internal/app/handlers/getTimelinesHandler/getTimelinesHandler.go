package getTimelinesHandler

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getTimelinesService"
	"net/http"
)

const RouteURL string = "/timelines"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := users.UserId(chi.URLParam(r, "user_id"))
	req := getTimelinesService.Request{UserId: userId}
	res, err := getTimelinesService.Execute(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}
