package core

import (
	"github.com/joho/godotenv"
	"os"
)

// Config -> environment variable settings
type Config struct {
	Production bool
	MongoURI   string
	Address    string
}

func address() string {
	port := os.Getenv("ADDRESS")
	if len(port) > 0 {
		return port
	}
	return ":8080"
}

// NewConfig -> get a configuarion instance
func load() *Config {
	godotenv.Load()
	return &Config{
		Production: os.Getenv("GO_ENV") == "production",
		MongoURI:   os.Getenv("MONGO_URI"),
		Address:    address(),
	}
}
