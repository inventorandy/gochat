package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"
)

// UpdateConversation updates a conversation in the database
func (r *ConversationRepository) UpdateConversation(conversation *types.Conversation) (*types.Conversation, error) {
	// Run the Query
	if err := r.db.Updates(conversation).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation
	return conversation, nil
}
