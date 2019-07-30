package db

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var PrimarySession Session

func NewSession(ctx context.Context) (*Session, error) {
	// connect to mongo server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	if client == nil || err != nil {
		return nil, errors.Wrap(err, "Error connecting to mongo server")
	}
	sess := Session{Client: client}
	return &sess, nil
}

func InitPrimarySession() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(getMongoURI()))
	if client == nil || err != nil {
		return errors.Wrap(err, "Error connecting to mongo server")
	}
	PrimarySession = Session{Client: client}
	return nil
}

func CleanPrimarySession() {
	PrimarySession.Close(context.Background())
}

func getMongoURI() string {
	mongoURI := os.Getenv("MONGO_URI")
	return mongoURI
}
