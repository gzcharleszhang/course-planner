package db

import (
	"context"
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
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
}

func (sess Session) Timelines() *mongo.Collection {
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("timelines")
}

func (sess Session) Courses() *mongo.Collection {
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("courses")
}

func (sess Session) CourseRecords() *mongo.Collection {
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("course-records")
}

func (sess Session) TermRecords() *mongo.Collection {
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("term-records")
}

func (sess Session) Plans() *mongo.Collection {
	return sess.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("plans")
}
