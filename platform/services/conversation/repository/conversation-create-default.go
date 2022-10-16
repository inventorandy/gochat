package repository

import (
	"fmt"
	"gochat/platform/internal/ptr"
	"gochat/platform/services/conversation/types"
	"log"

	"github.com/google/uuid"
)

// CreateDefaultConversations creates the default channels for the chat application
func (r *ConversationRepository) CreateDefaultConversations() error {
	// Create a new UUID for `#general` channel
	generalID := uuid.New()
	log.Println("general ID: ", generalID.String())

	// Check if we have a `#general` channel
	general, _ := r.GetConversationByLabel("general")
	if general == nil {
		// Create the Conversation
		generalConversation := &types.Conversation{
			ID:       generalID,
			Label:    "general",
			IsPublic: ptr.Bool(true),
		}
		conv, err := r.CreateConversation(generalConversation)
		if err != nil {
			return fmt.Errorf("startup error: %s", err.Error())
		}
		log.Printf("Conversation: %v\r\n", conv)
	}

	return nil
}
