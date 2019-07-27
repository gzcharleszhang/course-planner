package utils

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/env"
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
