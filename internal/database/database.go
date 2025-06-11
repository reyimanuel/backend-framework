package database

import (
	"backend/config"
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB initializes and returns a GORM database connection and its underlying SQL connection with retry.
func ConnectDB() (*gorm.DB, *sql.DB) {
	cfg := config.Get() // Get the database configuration from the config package.

	// Set up a custom SQL logger to log queries
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		})

	log.Println("Connecting to database...")

	var db *gorm.DB
	var err error

	maxRetries := 5
	retryDelay := 5 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(postgres.Open(cfg.DbUri), &gorm.Config{
			Logger:                 sqlLogger,
			SkipDefaultTransaction: true,
			AllowGlobalUpdate:      false,
		})

		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}

		log.Printf("Failed to connect to the database (attempt %d/%d): %v", attempt, maxRetries, err)
		if attempt < maxRetries {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		} else {
			log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error obtaining SQL DB from GORM: %v", err)
	}

	// Set DB connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection pool configured.")
	return db, sqlDB
}
