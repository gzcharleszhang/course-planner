package loginHandler

import (
	"encoding/json"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/loginService"
	"net/http"
)

const RouteURL string = "/login"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	var req loginService.Request
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), 500)
		return
	}
	res, err := loginService.Execute(ctx, req)
	if err != nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}
	auth.SetJwtCookie(res.JWTToken, w)
	w.Write([]byte(utils.ToJson(res)))
}
