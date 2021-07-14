package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConversationRepository struct
type ConversationRepository struct {
	db *gorm.DB
}

// NewConversationRepository creates a new instance of the Conversation Repository
func NewConversationRepository() (*ConversationRepository, error) {
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
	dbData, exists := os.LookupEnv("CONVERSATION_DB")
	if !exists {
		return nil, fmt.Errorf("repository error: no environment variable for CONVERSATION_DB")
	}

	// Build the Connection String
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPass, dbData, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Migrate the Table Design
	if err := db.AutoMigrate(&types.Conversation{}); err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Migrate the Table Design
	if err := db.AutoMigrate(&types.ConversationHasParticipant{}); err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Migrate the Table Design
	if err := db.AutoMigrate(&types.Message{}); err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Create the Repository Object
	repo := &ConversationRepository{
		db: db,
	}

	// Create the Default Conversations
	log.Println("CREATING DEFAULT CONVERSATIONS")
	if err := repo.CreateDefaultConversations(); err != nil {
		return nil, err
	}

	// Return the Repository Object
	return repo, nil
}
