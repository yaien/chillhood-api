package migrations

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

type Migration struct {
	Name      string
	Upgrade   func(db *mongo.Database) error
	Downgrade func(db *mongo.Database) error
}

type Updater struct {
	CollectionName    string
	InitialSchemaName string
	initial           func(db *mongo.Database) error
	migrations        []*Migration
	db                *mongo.Database
}

func (u *Updater) collection() *mongo.Collection {
	return u.db.Collection(u.CollectionName)
}

func (u *Updater) checkInitial(ctx context.Context) (updated bool, err error) {
	if u.initial == nil {
		return false, nil
	}

	c := u.collection()
	err = c.FindOne(ctx, bson.M{"name": u.InitialSchemaName}).Err()
	if err == nil {
		return false, nil
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return false, fmt.Errorf("failed finding initial schema: %w", err)
	}

	err = u.initial(u.db)
	if err != nil {
		return false, fmt.Errorf("failed initializing schema: %w", err)
	}

	docs := make([]interface{}, len(u.migrations)+1)
	docs[0] = bson.M{"name": u.InitialSchemaName}
	for index, m := range u.migrations {
		docs[index+1] = bson.M{"name": m.Name}
	}

	_, err = c.InsertMany(ctx, docs)
	if err != nil {
		return false, fmt.Errorf("failed saving initial schema: %w", err)
	}

	return true, nil
}

func (u *Updater) migrate(ctx context.Context) error {
	sort.Slice(u.migrations, func(i, j int) bool {
		return u.migrations[i].Name < u.migrations[j].Name
	})

	c := u.collection()
	for _, mig := range u.migrations {
		err := c.FindOne(ctx, bson.M{"name": mig.Name}).Err()
		if err == nil {
			continue
		}
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("failed finding migration %s: %w", mig.Name, err)
		}

		err = mig.Upgrade(u.db)
		if err != nil {
			return fmt.Errorf("failed executing migration %s: %w", mig.Name, err)
		}

		_, err = c.InsertOne(ctx, bson.M{"name": mig.Name})
		if err != nil {
			return fmt.Errorf("failed sabing migration %s: %w", mig.Name, err)
		}
	}

	return nil
}

func (u *Updater) Update(ctx context.Context) error {
	updated, err := u.checkInitial(ctx)
	if err != nil {
		return fmt.Errorf("failed checking initial: %w", err)
	}
	if updated {
		return nil
	}

	return u.migrate(ctx)
}

func NewUpdater(db *mongo.Database) *Updater {
	return &Updater{
		CollectionName:    "migrations",
		InitialSchemaName: "INITIAL_SCHEMA",
		initial:           initial,
		migrations:        migrations,
		db:                db,
	}
}
