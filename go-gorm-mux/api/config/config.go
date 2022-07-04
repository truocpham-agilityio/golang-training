package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the server configuration
// that includes database configuration and server port.
type Config struct {
	DB   *DBConfig
	Port string
}

// DBConfig represents the database configuration.
type DBConfig struct {
	DbDialect      string
	DbUsername     string
	DbPassword     string
	DbPort         string
	DbHost         string
	DbName         string
	TestDbDialect  string
	TestDbUsername string
	TestDbPassword string
	TestDbPort     string
	TestDbHost     string
	TestDbName     string
}

// GetConfig returns the server configuration.
func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println("We are getting the env values")
	}

	return &Config{
		DB: &DBConfig{
			DbDialect:      os.Getenv("DB_DRIVER"),
			DbUsername:     os.Getenv("DB_USER"),
			DbPassword:     os.Getenv("DB_PASSWORD"),
			DbPort:         os.Getenv("DB_PORT"),
			DbHost:         os.Getenv("DB_HOST"),
			DbName:         os.Getenv("DB_NAME"),
			TestDbDialect:  os.Getenv("TEST_DB_DRIVER"),
			TestDbUsername: os.Getenv("TEST_DB_USER"),
			TestDbPassword: os.Getenv("TEST_DB_PASSWORD"),
			TestDbPort:     os.Getenv("TEST_DB_PORT"),
			TestDbHost:     os.Getenv("TEST_DB_HOST"),
			TestDbName:     os.Getenv("TEST_DB_NAME"),
		},
		Port: os.Getenv("PORT"),
	}
}
