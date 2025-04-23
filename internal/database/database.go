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

// ConnectDB initializes and returns a GORM database connection and its underlying SQL connection.
func ConnectDB() (*gorm.DB, *sql.DB) {
	cfg := config.Get() // Get the database configuration from the config package.

	// Set up a custom SQL logger to log queries
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Logs output to console
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		})

	log.Println("connecting to databases...")

	db, err := gorm.Open(postgres.Open(cfg.DbUri), &gorm.Config{ // Open connection to database using GORM and PostgreSQL
		Logger:                 sqlLogger, // Use the custom SQL logger
		SkipDefaultTransaction: true,      // Skip default transaction
		AllowGlobalUpdate:      false,     // Do not allow global updates to avoid accidental data loss
	})

	// Check if there's an error while connecting to the database
	if err != nil {
		log.Fatalf("error connect sql. error : %v", err)
	}

	log.Println("Set database connection configuration...")

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Error set database connection configuration. error : %v", err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, sqlDB
}
