package infrastructure

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/yaien/p2p"
	"go.mongodb.org/mongo-driver/mongo"
)

// App -> application dependencies
type App struct {
	DB        *mongo.Database
	Config    *Config
	Slack     *slack.Client
	Templates *Templates
	P2P       *p2p.P2P
}

// NewApp return new application instance
func NewApp() (*App, error) {
	config := load()

	templates, err := parseTemplates(config)
	if err != nil {
		return nil, fmt.Errorf("failed parsing templates: %w", err)
	}

	slackClient := slack.New(config.Slack.AccessToken)

	db, err := connect(config.MongoURI, config.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the database: %w", err)
	}

	p := p2p.New(p2p.Options{
		Name:      config.P2P.Name,
		Addr:      config.Address,
		Lookup:    config.P2P.Lookup,
		Transport: &p2p.HttpTransport{Key: config.P2P.Key},
	})

	app := &App{
		DB:        db,
		Config:    config,
		Slack:     slackClient,
		Templates: templates,
		P2P:       p,
	}

	return app, nil
}
