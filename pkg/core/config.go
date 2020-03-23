package core

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type EpaycoConfig struct {
	CustomerID string
	Key        string
	PublicKey  string
	Test       bool
}

type JWTConfig struct {
	Secret   []byte
	Duration time.Duration
}

type ClientConfig struct {
	Keys    []string
	Origins []string
}

// Config -> environment variable settings
type Config struct {
	Production bool
	MongoURI   string
	Address    string
	BaseURL    *url.URL
	Epayco     *EpaycoConfig
	JWT        *JWTConfig
	Client     *ClientConfig
}

func address() string {
	port := os.Getenv("PORT")
	if len(port) > 0 {
		return fmt.Sprintf(":%s", port)
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

func baseURL(env string) *url.URL {
	u, err := url.Parse(env)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

// NewConfig -> get a configuration instance
func load() *Config {
	godotenv.Load()
	return &Config{
		Production: os.Getenv("GO_ENV") == "production",
		MongoURI:   os.Getenv("MONGODB_URI"),
		Address:    address(),
		BaseURL:    baseURL(os.Getenv("BASE_URL")),
		Epayco: &EpaycoConfig{
			Key:        os.Getenv("EPAYCO_KEY"),
			CustomerID: os.Getenv("EPAYCO_CUSTOMER_ID"),
			PublicKey:  os.Getenv("EPAYCO_PUBLIC_KEY"),
			Test:       os.Getenv("EPAYCO_TEST_MODE") != "false",
		},
		JWT: &JWTConfig{
			Secret:   []byte(os.Getenv("JWT_SECRET")),
			Duration: expiration(os.Getenv("JWT_DURATION")),
		},
		Client: &ClientConfig{
			Keys:    strings.Split(os.Getenv("CLIENT_KEY"), ","),
			Origins: strings.Split(os.Getenv("CLIENT_ORIGIN"), ","),
		},
	}
}
