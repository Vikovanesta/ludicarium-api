package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port         string
	DBDriver     string
	DBUri        string
	ClientID     string
	ClientSecret string
}

func EnvConfig() Config {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "3000"
	}

	dbDriver, ok := os.LookupEnv("DB_DRIVER")

	if !ok {
		panic("DB_DRIVER is not set")
	}

	dbURI, ok := os.LookupEnv("DB_URI")

	if !ok {
		panic("DB_URI is not set")
	}

	clientID, ok := os.LookupEnv("IGDB_CLIENT_ID")

	if !ok {
		panic("IGDB_CLIENT_ID is not set")
	}

	clientSecret, ok := os.LookupEnv("IGDB_CLIENT_SECRET")

	if !ok {
		panic("IGDB_CLIENT_SECRET is not set")
	}

	return Config{port, dbDriver, dbURI, clientID, clientSecret}
}
