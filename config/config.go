package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port         string
	DBUri        string
	ClientID     string
	ClientSecret string
}

func EnvConfig() Config {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "3000"
	}

	dbURI, ok := os.LookupEnv("POSTGRESQL_URL")

	if !ok {
		panic("POSTGRESQL_URL is not set")
	}

	clientID, ok := os.LookupEnv("IGDB_CLIENT_ID")

	if !ok {
		panic("IGDB_CLIENT_ID is not set")
	}

	clientSecret, ok := os.LookupEnv("IGDB_CLIENT_SECRET")

	if !ok {
		panic("IGDB_CLIENT_SECRET is not set")
	}

	return Config{port, dbURI, clientID, clientSecret}
}
