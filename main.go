package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vikovanesta/ludicarium-api/config"
	"github.com/vikovanesta/ludicarium-api/server"
)

func main() {
	config := config.EnvConfig()

	server := server.NewServer()

	log.Fatal(server.Run(config.Port))
}
