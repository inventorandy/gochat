package repository

import (
	"fmt"
	"gochat/platform/services/account/interfaces"
	"gochat/platform/services/account/types"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// AccountRepository struct
type AccountRepository struct {
	interfaces.Repository
	db *gorm.DB
}

// NewAccountRepository creates a new instance of the Account Repository
func NewAccountRepository() (*AccountRepository, error) {
	// Get the Database Credentials
	dbHost, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for POSTGRES_HOST")
	}
	dbPort, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for POSTGRES_PORT")
	}
	dbUser, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for POSTGRES_USER")
	}
	dbPass, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for POSTGRES_PASSWORD")
	}
	dbData, exists := os.LookupEnv("ACCOUNT_DB")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for ACCOUNT_DB")
	}

	// Build the Connection String
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPass, dbData, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Migrate the Table Design
	if err := db.AutoMigrate(&types.User{}); err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Create the Repository Object
	repo := &AccountRepository{
		db: db,
	}

	// Return the Repository Object
	return repo, nil
}
