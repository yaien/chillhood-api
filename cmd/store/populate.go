package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/assets"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func populate() *cobra.Command {
	return &cobra.Command{
		Use:   "db:populate",
		Short: "update initial seeder's data into db",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := core.NewApp()
			if err != nil {
				return err
			}
			if err := populateCities(app.DB); err != nil {
				return err
			}
			return nil
		},
	}
}

func populateCities(db *mongo.Database) error {
	var cities []struct {
		Days     int    `yaml:"days"`
		Name     string `yaml:"name"`
		Province string `yaml:"province"`
		Shipment int    `yaml:"shipment"`
	}
	source, err := assets.FS().ReadFile("seeders/cities.yaml")
	if err != nil {
		return fmt.Errorf("failed reading seeder file: %w", err)
	}

	if err := yaml.Unmarshal(source, &cities); err != nil {
		return err
	}
	provinces := make(map[string]models.Province)
	for _, city := range cities {
		province, ok := provinces[city.Province]
		if !ok {
			collection := db.Collection("provinces")
			filter := bson.M{"name": city.Province}
			err := collection.FindOne(context.TODO(), filter).Decode(&province)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					province = models.Province{ID: primitive.NewObjectID(), Name: city.Province}
					if _, err := collection.InsertOne(context.TODO(), province); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("failed finding province %s: %w", city.Province, err)
				}
			}
			provinces[city.Province] = province
		}
		collection := db.Collection("cities")
		filter := bson.M{"name": city.Name, "province._id": province.ID}
		created := false
		cty := models.City{}
		err = collection.FindOne(context.TODO(), filter).Decode(&cty)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				cty = models.City{
					ID:       primitive.NewObjectID(),
					Name:     city.Name,
					Days:     city.Days,
					Shipment: city.Shipment,
					Province: &province,
				}
				_, err := collection.InsertOne(context.TODO(), &cty)
				if err != nil {
					return err
				}
				created = true
			} else {
				return fmt.Errorf("failed finding city %s: %w", city.Name, err)
			}
		}
		if created {
			continue
		}
		_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": cty.ID}, bson.M{
			"$set": bson.M{
				"shipment": city.Shipment,
				"days":     city.Days,
				"province": &province,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
