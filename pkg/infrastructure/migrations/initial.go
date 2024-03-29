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
			if err == nil {
				provinces[city.Province] = *pr
				province = *pr
				continue
			}

			var ce *entity.Error
			if !errors.As(err, &ce) || ce.Code != "NOT_FOUND" {
				return fmt.Errorf("failed finding province: %w", err)
			}

			pr = &entity.Province{Name: city.Province}
			err = repo.Create(context.TODO(), pr)
			if err != nil {
				return fmt.Errorf("failed creating province: %w", err)
			}
			provinces[city.Province] = *pr
			province = *pr
		}

		repo := mongodb.NewCityRepository(db)
		rc, err := repo.FindOne(context.TODO(), entity.FindOneCityOptions{
			Name:       city.Name,
			ProvinceID: province.ID,
		})

		if err == nil {
			rc.Days = city.Days
			rc.Shipment = city.Shipment
			rc.Province = &province
			err = repo.Update(context.TODO(), rc)
			if err != nil {
				return fmt.Errorf("failed updating city: %w", err)
			}
			continue
		}

		var me *entity.Error
		if !errors.As(err, &me) || me.Code != "NOT_FOUND" {
			return fmt.Errorf("failed finding city: %w", err)
		}

		rc = &entity.City{
			Name:     city.Name,
			Days:     city.Days,
			Shipment: city.Shipment,
			Province: &province,
		}

		err = repo.Create(context.TODO(), rc)
		if err != nil {
			return fmt.Errorf("failed creating city: %w", err)
		}

	}
	return nil
}

func populateUsers(db *mongo.Database) error {
	users := []*entity.User{
		{
			Name:     "CODE",
			Email:    "stevensonxmarquez@gmail.com",
			Phone:    "+573163235111",
			Password: "nomelase",
		},
		{
			Name:     "FEAR",
			Email:    "felipechr14@gmail.com",
			Phone:    "+573014700584",
			Password: "nomelase",
		},
	}

	docs := make([]interface{}, len(users))
	for i, user := range users {
		user.ID = primitive.NewObjectID()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Role = "admin"
		_ = user.HashPassword()
		docs[i] = user
	}

	_, err := db.Collection("users").InsertMany(context.TODO(), docs)
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
