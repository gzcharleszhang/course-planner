package newUserService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/pkg/errors"
)

type Request struct {
	FirstName users.FirstName `json:"first_name"`
	LastName  users.LastName  `json:"last_name"`
	Password  string          `json:"password"`
}

type Response struct {
	UserId users.UserId `json:"user_id"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	hash, err := users.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "Error hashing user password")
	}
	userId, err := users.CreateUser(ctx, req.FirstName, req.LastName, hash)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new user")
	}

	return &Response{UserId: userId}, nil
}
