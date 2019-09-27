package scripts

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/env"
	"os"
	"time"
)

func InitScript() (context.Context, error) {
	env.LoadEnv()
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Minute)
	err := db.InitPrimarySession()
	if err != nil {
		return ctx, err
	}
	return ctx, db.PrimarySession.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Drop(ctx)
}

func CleanUpScript() {
	db.ClosePrimarySession()
}
