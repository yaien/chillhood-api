package core

import (
	"github.com/joho/godotenv"
	"os"
)

type EpaycoConfig struct {
	CustomerID string
	Key        string
	PublicKey  string
	Test       bool
}

// Config -> environment variable settings
type Config struct {
	Production bool
	MongoURI   string
	Address    string
	Epayco     EpaycoConfig
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
		Epayco: EpaycoConfig{
			Key:        os.Getenv("EPAYCO_KEY"),
			CustomerID: os.Getenv("EPAYCO_CUSTOMER_ID"),
			PublicKey:  os.Getenv("EPAYCO_PUBLIC_KEY"),
			Test:       os.Getenv("EPAYCO_TEST_MODE") != "false",
		},
	}
}
