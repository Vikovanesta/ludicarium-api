package db

import (
	"log"
	"os"
	"time"

	"github.com/vikovanesta/ludicarium-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 50, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	config := config.EnvConfig()

	uri := config.DBUri
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting *sql.DB instance:", err)
		return nil, err
	}

	// Ping the database
	err = pgsqlDB.Ping()
	if err != nil {
		log.Println("Error pinging database:", err)
		return nil, err
	}

	pgsqlDB.SetMaxIdleConns(10)
	pgsqlDB.SetMaxOpenConns(100)
	pgsqlDB.SetConnMaxLifetime(time.Hour)

	return &DB{db}, nil
}
