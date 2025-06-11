package config

import (
	"backend/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// This is a struct that holds the configuration of the application.
type AppConfigurationMap struct {
	Port                 int    // Port is the port number that the server will listen to.
	IsProduction         bool   // IsProduction is a flag that indicates whether the application is running in production mode.
	DbUri                string // Database connection.
	AccessTokenLifeTime  uint   // AccessTokenLifeTime is the lifetime of the access token in seconds.
	RefreshTokenLifeTime uint   // RefreshTokenLifeTime is the lifetime of the refresh token in seconds.
	PrivateKeyPEM        string // was PrivateKeyPath
	PublicKeyPEM         string
	BaseURL              string // BaseURL is the base URL of the application, used for generating absolute URLs.
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

	// Load environment variables from a .env file.
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables, try to get from environtment OS")
	}

	// Read the PORT environment variable and convert it to an integer.
	port, err := strconv.Atoi(os.Getenv("PORT"))

	// Set default value port if env doesn't have PORT config
	if err != nil {
		port = 8080
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

	privateKeyPEM := os.Getenv("PRIVATE_KEY")
	if privateKeyPEM == "" {
		log.Fatalf("PRIVATE_KEY environment variable is not set, check your Railway environment variables")
	}
	
	publicKeyPEM := os.Getenv("PUBLIC_KEY")
	if publicKeyPEM == "" {
		log.Fatalf("PUBLIC_KEY environment variable is not set, check your Railway environment variables")
	}
	
	// Replace literal '\n' from Railway's env to actual newlines
	privateKeyPEM = strings.ReplaceAll(privateKeyPEM, "\\n", "\n")
	publicKeyPEM = strings.ReplaceAll(publicKeyPEM, "\\n", "\n")
	
	BaseURL := os.Getenv("BASE_URL")
	if BaseURL == "" {
		BaseURL = fmt.Sprintf("http://localhost:%d", port)
	}

	// Set global variable config
	config = &AppConfigurationMap{
		Port:                 port,
		IsProduction:         isProduction,
		DbUri:                loadDatabaseConfig(),
		AccessTokenLifeTime:  uint(AccessTokenLifeTime),
		RefreshTokenLifeTime: uint(RefreshTokenLifeTime),
		PrivateKeyPEM:        privateKeyPEM, 
		PublicKeyPEM:         publicKeyPEM,
		BaseURL:              BaseURL,
	}
}

// loadDatabaseConfig is a function that loads the database configuration from the environment variables.
func loadDatabaseConfig() string {
	user := getFromEnv("DB_USER")
	pass := getFromEnv("DB_PASS")
	name := getFromEnv("DB_NAME")
	host := getFromEnv("DB_HOST")
	port := getFromEnv("DB_PORT")
	timeZone := getFromEnv("DB_TIME_ZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", host, user, pass, name, port, timeZone)
}

// getFromEnv retrieves an environment variable by key and exits the program if it's not set.
func getFromEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		// If the environment variable is missing, log an error and terminate the application.
		log.Fatalf("%s Environment variable is not set", value)
	}

	return value
}
