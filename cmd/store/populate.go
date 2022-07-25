package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/assets"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func populate() *cobra.Command {
	return &cobra.Command{
		Use:   "db:populate",
		Short: "update initial seeder's data into db",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := infrastructure.NewApp()
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

	provinces := make(map[string]entity.Province)
	for _, city := range cities {
		province, ok := provinces[city.Province]
		if !ok {
			repo := mongodb.NewProvinceRepository(db)
			pr, err := repo.FindOneByName(context.TODO(), city.Province)
			if err != nil {
				var ce *entity.Error
				if errors.As(err, &ce) && ce.Code == "not_found" {
					pr = &entity.Province{Name: city.Province}
					err := repo.Create(context.TODO(), pr)
					if err != nil {
						return fmt.Errorf("failed creating province: %w", err)
					}
				} else {
					return err
				}
			}
			provinces[city.Province] = *pr
			province = *pr
		}

		repo := mongodb.NewCityRepository(db)
		var created bool
		city, err := repo.FindOne(context.TODO(), entity.FindOneCityOptions{
			Name:       city.Name,
			ProvinceID: province.ID,
		})
		if err != nil {
			var me *entity.Error
			if errors.As(err, &me) && me.Code == "not_found" {
				city = &entity.City{
					Name:     city.Name,
					Days:     city.Days,
					Shipment: city.Shipment,
					Province: &province,
				}
				err := repo.Create(context.TODO(), city)
				if err != nil {
					return err
				}
				created = true
			} else {
				return err
			}
		}

		if created {
			continue
		}
		err = repo.Update(context.TODO(), city)
		if err != nil {
			return err
		}
	}
	return nil
}
