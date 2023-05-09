package mongo_store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseName string

type DatabaseURI string

type Options struct {
	DatabaseName DatabaseName
	Client       *mongo.Client
}

var ErrDBNameNotSpecified = errors.New("database name is not specified in database URI")

// New returns either error or mongo client alongside db name
func New(uri DatabaseURI) (Options, error) {
	dbURI := string(uri)

	client, err := connectToDB(dbURI)
	if err != nil {
		return Options{}, err
	}

	dbNameStartIndex := strings.LastIndex(dbURI, "/")
	dbNameEndIndex := strings.LastIndex(dbURI, "?")
	if dbNameStartIndex < 0 || dbNameEndIndex < 0 {
		return Options{}, ErrDBNameNotSpecified
	}

	dbName := dbURI[dbNameStartIndex+1 : dbNameEndIndex]

	return Options{
		DatabaseName: DatabaseName(dbName),
		Client:       client,
	}, nil
}

func connectToDB(dbURI string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(dbURI)
	timeout := time.Second * 30
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo: failed to connect: %w", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo: failed to ping: %w", err)
	}

	return client, nil
}
