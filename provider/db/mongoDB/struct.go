package gpmongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDB(ctx context.Context, clientOptions *ClientOptions) (*MongoDB, error) {

	opts := options.Client().
		ApplyURI(clientOptions.URI).
		SetConnectTimeout(time.Duration(clientOptions.ConnectTimeout) * time.Second).
		SetMaxConnIdleTime(time.Duration(clientOptions.MaxConnIdleTime) * time.Second).
		SetMaxPoolSize(clientOptions.MaxPoolSize).
		SetMinPoolSize(clientOptions.MaxPoolSize).
		SetServerSelectionTimeout(time.Duration(clientOptions.ServerSelectionTimeout) * time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect failed: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("client.Ping failed: %w", err)
	}

	database := client.Database(clientOptions.DB)

	mongoDB := MongoDB{
		Client:   client,
		Database: database,
	}

	return &mongoDB, nil
}

func (that *MongoDB) SetCollection(collection string) {

	that.Database.Collection(collection)
}

func (that *MongoDB) Disconnect(ctx context.Context) error {

	if err := that.Client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func (that *MongoDB) BulkWrite(
	ctx context.Context,
	writes []mongo.WriteModel,
	retries int,
	timeout time.Duration,
	opts *options.BulkWriteOptions,
) error {

	var err error

	for i := 0; i < retries; i++ {
		ctxWrite, cancel := context.WithTimeout(ctx, timeout)

		_, err = that.Collection.BulkWrite(ctxWrite, writes, opts)

		cancel()

		if err == nil {
			return nil
		}

		time.Sleep(100 * time.Millisecond)
	}

	return fmt.Errorf("bulkInsert failed: %w", err)
}
