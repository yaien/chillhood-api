package infrastructure

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

type CloudinaryConfig struct {
	CloudName    string
	UploadPreset string
}

type SlackConfig struct {
	Channel     string
	AccessToken string
	SaleUrl     string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Sender   string
	RefLink  string
}

func (s *SMTPConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Config -> environment variable settings
type Config struct {
	Production bool
	MongoURI   string
	MongoDB    string
	Address    string
	BaseURL    *url.URL
	Epayco     *EpaycoConfig
	JWT        *JWTConfig
	Client     *ClientConfig
	Cloudinary *CloudinaryConfig
	Slack      *SlackConfig
	SMTP       *SMTPConfig
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
		MongoURI:   os.Getenv("MONGO_URL"),
		MongoDB:    os.Getenv("MONGO_DB"),
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
			Keys:    strings.Split(os.Getenv("CLIENT_KEYS"), ","),
			Origins: strings.Split(os.Getenv("CLIENT_ORIGIN"), ","),
		},
		Cloudinary: &CloudinaryConfig{
			CloudName:    os.Getenv("CLOUDINARY_CLOUD_NAME"),
			UploadPreset: os.Getenv("CLOUDINARY_UPLOAD_PRESET"),
		},
		Slack: &SlackConfig{
			AccessToken: os.Getenv("SLACK_ACCESS_TOKEN"),
			Channel:     os.Getenv("SLACK_CHANNEL"),
			SaleUrl:     os.Getenv("SLACK_SALE_URL"),
		},
		SMTP: &SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Sender:   os.Getenv("SMTP_SENDER"),
			RefLink:  os.Getenv("SMTP_REF_LINK"),
		},
	}
}
