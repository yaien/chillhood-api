package core

import (
	"fmt"
	"github.com/slack-go/slack"
	"go.mongodb.org/mongo-driver/mongo"
)

// App -> application dependencies
type App struct {
	DB        *mongo.Database
	Config    *Config
	Slack     *slack.Client
	Templates *Templates
}

// NewApp return new application instance
func NewApp() (*App, error) {
	config := load()
	db, err := connect(config.MongoURI)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the database: %w", err)
	}

	templates, err := parseTemplates()
	if err != nil {
		return nil, fmt.Errorf("failed parsing templates: %w", err)
	}

	slackClient := slack.New(config.Slack.AccessToken)

	app := &App{
		DB:        db,
		Config:    config,
		Slack:     slackClient,
		Templates: templates,
	}

	return app, nil
}
