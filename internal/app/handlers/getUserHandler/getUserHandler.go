package getUserHandler

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getUserService"
	"net/http"
)

const RouteURL string = "/user/{user_id}"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := users.UserId(chi.URLParam(r, "user_id"))
	// for regular users, they can only access their own user data
	if roles.GetRoleFromContext(ctx).GetRoleId() == roles.ConradId &&
		users.GetUserIdFromContext(ctx) != userId {
		http.Error(w, http.StatusText(401), 401)
		return
	}
	req := getUserService.Request{UserId: userId}
	res, err := getUserService.Execute(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(utils.ToJson(res)))
}
