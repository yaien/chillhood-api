package core

import "go.mongodb.org/mongo-driver/mongo"

// App -> application dependencies
type App struct {
	DB     *mongo.Database
	Config *Config
}

// NewApp return new application instance
func NewApp() (*App, error) {
	config := load()
	db, err := connect(config.MongoURI)
	if err != nil {
		return nil, err
	}
	return &App{db, config}, nil
}
