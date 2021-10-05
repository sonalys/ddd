package persistence

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "cart-api"

// Mongo is an adapter for persistence layer.
type Mongo struct {
	*mongo.Database
}

// MongoConfig are the configurations required to instantiate the mongo adapter.
type MongoConfig struct {
	Hosts    []string `json:"hosts"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

// NewMongoClient instantiates a new mongo client.
func NewMongoClient(ctx context.Context, c MongoConfig) (*Mongo, error) {
	mongoConfig := options.Client()
	mongoConfig.Hosts = c.Hosts
	mongoConfig.Auth = &options.Credential{
		Username: c.Username,
		Password: c.Password,
	}

	col, err := mongo.NewClient(mongoConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize persistence client")
	}

	db := col.Database(databaseName)

	return &Mongo{
		Database: db,
	}, nil
}
