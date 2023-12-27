package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vikovanesta/ludicarium-api/server"
)

type config struct {
	port  string
	dbUri string
}

func main() {
	config := envConfig()

	server := server.NewServer()

	log.Fatal(server.Run(config.port))
}

func envConfig() config {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "3000"
	}

	dbURI, ok := os.LookupEnv("POSTGRESQL_URL")

	if !ok {
		panic("POSTGRESQL_URL is not set")
	}

	return config{port, dbURI}
}
