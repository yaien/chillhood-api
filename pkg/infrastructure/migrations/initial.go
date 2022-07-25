package migrations

import (
	"context"
	"errors"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/assets"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
	"time"
)

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

	err = yaml.Unmarshal(source, &cities)
	if err != nil {
		return fmt.Errorf("failed unmarshall cities.yaml")
	}

	provinces := make(map[string]entity.Province)
	for _, city := range cities {
		province, ok := provinces[city.Province]
		if !ok {
			repo := mongodb.NewProvinceRepository(db)
			pr, err := repo.FindOneByName(context.TODO(), city.Province)
			if err != nil {
				var ce *entity.Error
				if errors.As(err, &ce) && ce.Code == "NOT_FOUND" {
					pr = &entity.Province{Name: city.Province}
					err := repo.Create(context.TODO(), pr)
					if err != nil {
						return fmt.Errorf("failed creating province: %w", err)
					}
				} else {
					return fmt.Errorf("failed finding province: %w", err)
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
			if errors.As(err, &me) && me.Code == "NOT_FOUND" {
				city = &entity.City{
					Name:     city.Name,
					Days:     city.Days,
					Shipment: city.Shipment,
					Province: &province,
				}
				err := repo.Create(context.TODO(), city)
				if err != nil {
					return fmt.Errorf("failed creating city: %w", err)
				}
				created = true
			} else {
				return fmt.Errorf("failed finding city: %w", err)
			}
		}

		if created {
			continue
		}
		err = repo.Update(context.TODO(), city)
		if err != nil {
			return fmt.Errorf("failed updating city: %w", err)
		}
	}
	return nil
}

func populateUsers(db *mongo.Database) error {
	users := []interface{}{
		&entity.User{
			Name:      "CODE",
			Email:     "stevensonxmarquez@gmail.com",
			Phone:     "+573163235111",
			Password:  "nomelase",
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		&entity.User{
			Name:      "FEAR",
			Email:     "felipechr14@gmail.com",
			Phone:     "+573014700584",
			Password:  "nomelase",
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := db.Collection("users").InsertMany(context.TODO(), users)
	return err
}

func concat(fs ...func(db *mongo.Database) error) func(db *mongo.Database) error {
	return func(db *mongo.Database) error {
		for _, f := range fs {
			err := f(db)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

var initial = concat(populateCities, populateUsers)

var migrations []*Migration

func register(mig *Migration) {
	migrations = append(migrations, mig)
}
