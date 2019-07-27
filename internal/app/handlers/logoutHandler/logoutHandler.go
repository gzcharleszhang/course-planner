package logoutHandler

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	auth.ClearJwtCookie(w)
	w.Write([]byte(utils.ToJson(utils.M{"message": "Logged out successfully."})))
}
