package db

import (
	"log"
	"os"
	"time"

	"github.com/vikovanesta/ludicarium-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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
	driver := config.DBDriver
	uri := config.DBUri

	if driver == "sqlite" {
		db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Println("Failed to connect to the database:", err)
			return nil, err
		}

		sqliteDB, err := db.DB()
		if err != nil {
			log.Println("Error getting *sql.DB instance:", err)
			return nil, err
		}

		// Ping the database
		err = sqliteDB.Ping()
		if err != nil {
			log.Println("Error pinging database:", err)
			return nil, err
		}

		sqliteDB.SetMaxIdleConns(10)
		sqliteDB.SetMaxOpenConns(100)
		sqliteDB.SetConnMaxLifetime(time.Hour)

		return &DB{db}, nil
	}

	if driver == "postgres" {
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

	if driver == "mysql" {
		db, err := gorm.Open(mysql.Open(uri), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Println("Failed to connect to the database:", err)
			return nil, err
		}

		mysqlDB, err := db.DB()
		if err != nil {
			log.Println("Error getting *sql.DB instance:", err)
			return nil, err
		}

		// Ping the database
		err = mysqlDB.Ping()
		if err != nil {
			log.Println("Error pinging database:", err)
			return nil, err
		}

		mysqlDB.SetMaxIdleConns(10)
		mysqlDB.SetMaxOpenConns(100)
		mysqlDB.SetConnMaxLifetime(time.Hour)

		return &DB{db}, nil
	}

	return nil, nil
}
