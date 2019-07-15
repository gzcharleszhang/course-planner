package userHandlers

import (
	"encoding/json"
	"fmt"
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
	if err = newUserService.Run(ctx, req); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), 500)
	}
}
