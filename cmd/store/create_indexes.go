package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndexes() *cobra.Command {
	return &cobra.Command{
		Use:   "db:create-indexes",
		Short: "create database indexes",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := core.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			for _, c := range collections {
				indexes := app.DB.Collection(c.name).Indexes()
				names, err := indexes.CreateMany(context.TODO(), c.indexes)
				if err != nil {
					fmt.Printf("create indexes failed on collection %s: %s\n", c.name, err)
					continue
				}
				fmt.Printf("created indexes %s on collection %s\n",
					strings.Join(names, ", "), c.name)
			}

		},
	}
}

var collections = []struct {
	name    string
	indexes []mongo.IndexModel
}{
	{
		name: "invoices",
		indexes: []mongo.IndexModel{
			{
				Keys: bson.M{
					"ref":            "text",
					"shipping.email": "text",
					"shipping.name":  "text",
					"shipping.phone": "text",
				},
				Options: options.Index().SetName("search_text"),
			},
			{Keys: bson.M{"ref": 1}, Options: options.Index().SetUnique(true)},
		},
	},
	{
		name: "items",
		indexes: []mongo.IndexModel{
			{Keys: bson.M{"name": 1}, Options: options.Index().SetUnique(true)},
			{Keys: bson.M{"slug": 1}, Options: options.Index().SetUnique(true)},
			{Keys: bson.M{"name": "text"}},
		},
	},
	{
		name: "users",
		indexes: []mongo.IndexModel{
			{
				Keys:    bson.M{"email": 1},
				Options: options.Index().SetUnique(true),
			},
		},
	},
}
