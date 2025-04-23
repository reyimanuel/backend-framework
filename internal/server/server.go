package server

import (
	"backend/config"
	"backend/controller"
	"backend/internal/database"
	"backend/migrations"
	"backend/repository"
	"backend/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Run is a function that starts the server and runs the database migrations.
// It also handles command line flags for running migrations down and down all.
// The function initializes the database connection, sets up the server, and starts listening for incoming requests.
// to run this server, use the command: go run main.go, but this project uses air to run the server automatically when there is a change in the code
// example: air
func Run() {
	log.Println("Starting server...")

	cfg := config.Get() // Get the application configuration
	if cfg == nil {
		log.Fatal("Failed to load configuration")
		return
	}

	db, sqlDB := database.ConnectDB() // Connect to the database
	if db == nil {
		log.Fatal("Failed to connect to the database")
		return
	}

	downFlag := flag.Bool("down", false, "Run database migration down")
	// Flag to run a single migration down. example: go run main.go down

	downAllFlag := flag.Bool("down-all", false, "Run all database migrations down")
	// Flag to run all migrations down. example: go run main.go down-all

	// Parse the flags
	flag.Parse()

	// If `--down` flag is passed, rollback the latest migration
	if *downFlag {
		log.Println("Running database migration down...")
		migrations.Down(sqlDB)
		log.Println("Successfully run database migration down.")
		return
	}

	// If `--down-all` flag is passed, rollback all migrations
	if *downAllFlag {
		log.Println("Running all database migrations down...")
		migrations.DownAll(sqlDB)
		log.Println("Successfully run all database migrations down.")
		return
	}

	// By default, apply all pending migrations
	log.Println("Running database migration up...")
	migrations.Up(sqlDB)
	log.Println("Successfully run database migration up.")

	// Initialize repositories with the connected GORM DB
	repo := repository.New(db)

	// Initialize service layer with repositories
	serv := service.New(repo)

	// Set Gin to production mode if the config indicates it
	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new Gin engine instance
	r := gin.New()

	// Register middleware for logging and error recovery
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set up HTTP server settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Set up routes and controllers
	controller.New(r, serv)

	log.Println("Server running successfully")

	// Start the server
	log.Printf("Server is running on port %d", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
