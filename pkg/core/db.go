package core

import (
	"context"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongo database
func connect(rawurl string) (*mongo.Database, error) {
	uri, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	opts := options.Client().ApplyURI(rawurl)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	path := strings.Replace(uri.Path, "/", "", 1)
	db := client.Database(path)
	return db, nil
}
