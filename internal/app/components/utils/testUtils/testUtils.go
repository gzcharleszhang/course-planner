package testUtils

import (
	"bytes"
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/env"
	"github.com/gzcharleszhang/course-planner/internal/app/services/getUserService"
	"github.com/gzcharleszhang/course-planner/internal/app/services/newUserService"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func Init() (context.Context, error) {
	env.LoadTestEnv()
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	err := db.InitPrimarySession()
	if err != nil {
		return ctx, err
	}
	return ctx, db.PrimarySession.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Drop(ctx)
}

func CleanUp() {
	db.CleanPrimarySession()
}

func InitWithUser() (context.Context, *users.User, error) {
	ctx, err := Init()
	if err != nil {
		return ctx, nil, err
	}
	req := newUserService.Request{
		FirstName: "Steven",
		LastName:  "Xu",
		Email:     "hello@stevenxu.me",
		Password:  "mrcalcaward",
	}
	res, err := newUserService.Execute(ctx, req)
	if err != nil {
		return ctx, nil, err
	}
	getReq := getUserService.Request{
		UserId: res.UserId,
	}
	getRes, err := getUserService.Execute(ctx, getReq)
	if err != nil {
		return ctx, nil, err
	}
	return ctx, &getRes.User, nil
}

func NewRequest(ctx context.Context, method, url string, requestBody []byte, handler http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr, nil
}
