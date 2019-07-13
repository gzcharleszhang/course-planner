package db

import "go.mongodb.org/mongo-driver/mongo"

type Session struct {
	Client *mongo.Client
}

func (sess Session) Users() *mongo.Collection {
	return sess.Client.Database("entities").Collection("users")
}

func (sess Session) Timelines() *mongo.Collection {
	return sess.Client.Database("entities").Collection("timelines")
}

func (sess Session) Courses() *mongo.Collection {
	return sess.Client.Database("entities").Collection("courses")
}

func (sess Session) CourseRecords() *mongo.Collection {
	return sess.Client.Database("entities").Collection("course-records")
}

func (sess Session) TermRecords() *mongo.Collection {
	return sess.Client.Database("entities").Collection("term-records")
}

func (sess Session) Plans() *mongo.Collection {
	return sess.Client.Database("entities").Collection("plans")
}