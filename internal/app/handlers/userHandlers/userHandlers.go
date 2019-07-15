package userHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newUserService"
	"net/http"
)

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	var req newUserService.Request
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	res, err := newUserService.Execute(ctx, req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), 500)
	}
	w.Write([]byte(utils.ToJson(res)))
}
