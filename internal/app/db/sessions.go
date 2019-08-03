package db

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/env"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Session struct {
	Client *mongo.Client
}

func (sess Session) Close(ctx context.Context) {
	sess.Client.Disconnect(ctx)
}

func (sess Session) Users() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("users")
}

func (sess Session) Timelines() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("timelines")
}

func (sess Session) Courses() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("courses")
}

func (sess Session) CourseRecords() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("course-records")
}

func (sess Session) TermRecords() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("term-records")
}

func (sess Session) Plans() *mongo.Collection {
	return sess.Client.Database(os.Getenv(env.MongoDBNameEnvKey)).Collection("plans")
}
