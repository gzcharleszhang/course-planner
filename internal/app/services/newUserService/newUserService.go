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
}

func Run(ctx context.Context, req Request) error {
	hash, err := users.HashPassword(req.Password)
	if err != nil {
		return errors.Wrap(err, "Error hashing user password")
	}
	if err := users.CreateUser(ctx, req.FirstName, req.LastName, hash); err != nil {
		return errors.Wrap(err, "Error creating new user")
	}
}
