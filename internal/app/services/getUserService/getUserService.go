package getUserService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
)

type Request struct {
	UserId users.UserId `json:"user_id"`
}

type Response struct {
	User users.User `json:"user"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	user, err := users.GetUserById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res := Response{User: *user}
	return &res, nil
}
