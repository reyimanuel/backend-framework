package config

import (
	"backend/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// This is a struct that holds the configuration of the application.
type AppConfigurationMap struct {
	Port               int    // Port is the port number that the server will listen to.
	IsProduction       bool   // IsProduction is a flag that indicates whether the application is running in production mode.
	DbUri              string // Database connection string (DSN).
	AccessTokenLifeTime  uint   // AccessTokenLifeTime is the lifetime of the access token in seconds.
	RefreshTokenLifeTime uint   // RefreshTokenLifeTime is the lifetime of the refresh token in seconds.
	PrivateKeyPEM     string // Path to the private key file.
	PublicKeyPEM      string // Path to the public key file.
	BaseURL            string // BaseURL is the base URL of the application, used for generating absolute URLs.
}

// config is a global variable that stores the loaded application configuration.
var config *AppConfigurationMap

// Get is a function that returns the loaded application configuration.
func Get() *AppConfigurationMap {
	return config
}

// Load is a function that loads the application configuration from the environment variables.
func Load() {
	log.Println("Loading config from environment...")

	// Load environment variables from a .env file for local development.
	// This will be gracefully ignored in production environments like Railway where no .env file exists.
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file, relying on OS environment variables.")
	}

	// Read the PORT environment variable and convert it to an integer.
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Set default value port if env doesn't have PORT config
	}

	// Check if the application is running in production mode.
	isProduction := utils.SafeCompareString(os.Getenv("IS_PRODUCTION"), "true")

	AccessTokenLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE_TIME"))
	if err != nil {
		AccessTokenLifeTime = 3600 // Default value of 1 hour
	}

	RefreshTokenLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE_TIME"))
	if err != nil {
		RefreshTokenLifeTime = 86400 // Default value of 24 hours
	}

	// --- LOGIKA DATABASE YANG DIPERBARUI ---
	// Prioritaskan penggunaan DATABASE_URL yang disediakan oleh Railway.
	dbUri := os.Getenv("DATABASE_URL")
	
	// Jika DATABASE_URL tidak ada, coba bangun DSN dari variabel PG* individual.
	// Ini membuat konfigurasi fleksibel untuk lokal dan produksi.
	if dbUri == "" {
		log.Println("DATABASE_URL not found, building DSN from individual PG* variables...")

		host := os.Getenv("PGHOST")
		user := os.Getenv("PGUSER")
		pass := os.Getenv("PGPASSWORD")
		name := os.Getenv("PGDATABASE")
		dbPort := os.Getenv("PGPORT")

		// Host adalah variabel paling penting. Jika tidak ada, koneksi tidak mungkin.
		if host == "" {
			log.Fatalf("Database configuration is incomplete. Set DATABASE_URL or individual PG* variables (e.g., PGHOST).")
		}
		
		// Gunakan port default postgres jika tidak dispesifikasikan
		if dbPort == "" {
			dbPort = "5432"
		}

		// Buat string DSN (Data Source Name) yang dimengerti oleh GORM.
		// sslmode=disable sering digunakan untuk koneksi di jaringan internal yang aman seperti Railway.
		dbUri = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, dbPort)
	}
	// --- AKHIR LOGIKA DATABASE ---


	PrivateKeyPEM := os.Getenv("PRIVATE_KEY")
	if PrivateKeyPEM == "" {
		log.Fatalf("PRIVATE_KEY environment variable is not set, check your Railway environment variables")
	}
	
	PublicKeyPEM := os.Getenv("PUBLIC_KEY")
	if PublicKeyPEM == "" {
		log.Fatalf("PUBLIC_KEY environment variable is not set, check your Railway environment variables")
	}

	BaseURL := os.Getenv("BASE_URL")
	if BaseURL == "" {
		BaseURL = fmt.Sprintf("http://localhost:%d", port)
	}

	// Set global variable config
	config = &AppConfigurationMap{
		Port:                 port,
		IsProduction:         isProduction,
		DbUri:                dbUri, // Gunakan dbUri yang sudah benar
		AccessTokenLifeTime:  uint(AccessTokenLifeTime),
		RefreshTokenLifeTime: uint(RefreshTokenLifeTime),
		PrivateKeyPEM:       PrivateKeyPEM,
		PublicKeyPEM:        PublicKeyPEM,
		BaseURL:              BaseURL,
	}
	
	log.Println("Configuration loaded successfully.")
}