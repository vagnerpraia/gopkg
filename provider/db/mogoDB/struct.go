package gpmongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, clientOptions *ClientOptions) (*mongo.Client, *mongo.Database, error) {

	opts := options.Client().ApplyURI(clientOptions.URI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("mongo.Connect failed: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("client.Ping failed: %w", err)
	}

	database := client.Database(clientOptions.DB)

	return client, database, nil
}
