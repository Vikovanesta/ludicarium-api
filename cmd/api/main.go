package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vikovanesta/ludicarium-api/config"
	"github.com/vikovanesta/ludicarium-api/db"
	"github.com/vikovanesta/ludicarium-api/igdb"
	"github.com/vikovanesta/ludicarium-api/server"
)

func main() {
	config := config.EnvConfig()

	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	igdbClient := igdb.NewClient(nil)

	server := server.NewServer(db, igdbClient)

	log.Fatal(server.Run(config.Port))
}
