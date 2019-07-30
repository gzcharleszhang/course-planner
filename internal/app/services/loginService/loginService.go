package loginService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/models/userModel"
)

type Request struct {
	Email    users.Email `json:"email"`
	Password string      `json:"password"`
}

type Response struct {
	JWTToken string `json:"jwt_token"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	user, err := userModel.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = users.VerifyPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	_, token, err := auth.GenerateTokenForUser(user.Id)
	res := Response{
		JWTToken: token,
	}
	return &res, nil
}
