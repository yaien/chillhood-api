package core

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type EpaycoConfig struct {
	CustomerID string
	Key        string
	PublicKey  string
	Test       bool
}

type JWTConfig struct {
	Secret []byte
	Duration time.Duration
}

// Config -> environment variable settings
type Config struct {
	Production bool
	MongoURI   string
	Address    string
	Epayco     EpaycoConfig
	JWT        JWTConfig
}

func address() string {
	port := os.Getenv("ADDRESS")
	if len(port) > 0 {
		return port
	}
	return ":8080"
}

func expiration(env string) time.Duration {
	value, err := time.ParseDuration(env)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// NewConfig -> get a configuration instance
func load() *Config {
	godotenv.Load()
	return &Config{
		Production: os.Getenv("GO_ENV") == "production",
		MongoURI:   os.Getenv("MONGO_URI"),
		Address:    address(),
		Epayco: EpaycoConfig{
			Key:        os.Getenv("EPAYCO_KEY"),
			CustomerID: os.Getenv("EPAYCO_CUSTOMER_ID"),
			PublicKey:  os.Getenv("EPAYCO_PUBLIC_KEY"),
			Test:       os.Getenv("EPAYCO_TEST_MODE") != "false",
		},
		JWT: JWTConfig{
			Secret:     []byte(os.Getenv("JWT_SECRET")),
			Duration: expiration(os.Getenv("JWT_DURATION")),
		},
	}
}
