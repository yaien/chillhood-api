package core

import (
	"github.com/slack-go/slack"
	"go.mongodb.org/mongo-driver/mongo"
)

// App -> application dependencies
type App struct {
	DB     *mongo.Database
	Config *Config
	Slack  *slack.Client
}

// NewApp return new application instance
func NewApp() (*App, error) {
	config := load()
	db, err := connect(config.MongoURI)
	if err != nil {
		return nil, err
	}

	slackClient := slack.New(config.Slack.AccessToken)

	app := &App{
		DB:     db,
		Config: config,
		Slack:  slackClient,
	}

	return app, nil
}
