package database

import (
	"context"
	"time"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func GetDatabaseInstance(serverConfig *config.ServerConfig) (*mongo.Client, error) {
	// create a mongo client
	client, err := mongo.Connect(options.Client().ApplyURI(serverConfig.MongoUri))

	if err != nil {
		return nil, err
	}

	// create a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// ping the client
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

func AccessCollection(serverConfig *config.ServerConfig, client *mongo.Client, collectionName string) *mongo.Collection {
	// access the database
	database := client.Database(serverConfig.DatabaseName)

	// access the collection
	collection := database.Collection(collectionName)

	return collection
}
