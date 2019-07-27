package db

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewSession(ctx context.Context) (*Session, error) {
	// connect to mongo server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	if client == nil || err != nil {
		return nil, errors.Wrap(err, "Error connecting to mongo server")
	}
	session := Session{Client: client}
	return &session, nil
}

func getMongoURI() string {
	mongoURI := os.Getenv("MONGO_URI")
	return mongoURI
}
