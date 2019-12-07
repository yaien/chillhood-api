package core

import (
	"context"
	"net/url"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongo database
func connect(rawurl string) (*mongo.Database, error) {
	uri, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	opts := options.Client().ApplyURI(rawurl)
	client, err := mongo.Connect(context.TODO(), opts)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	path := strings.Replace(uri.Path, "/", "", 1)
	db := client.Database(path)
	return db, nil
}
