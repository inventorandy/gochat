package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"
)

// CreateConversation creates a new conversation in the database
func (r *ConversationRepository) CreateConversation(conversation *types.Conversation) (*types.Conversation, error) {
	// Run the Query
	if err := r.db.Create(conversation).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation
	return conversation, nil
}
