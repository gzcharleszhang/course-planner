package utils

import (
	"bytes"
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/env"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func InitTest() (context.Context, error) {
	env.LoadTestEnv()
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	sess, err := db.NewSession(ctx)
	if err != nil {
		return ctx, err
	}
	defer sess.Close(ctx)
	return ctx, sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Drop(ctx)
}

func NewTestRequest(method, url string, requestBody []byte, handler http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr, nil
}
