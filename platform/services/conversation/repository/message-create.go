package repository

import (
	"fmt"
	"gochat/platform/services/conversation/types"
)

// CreateMessage creates a new message in the database
func (r *ConversationRepository) CreateMessage(message *types.Message) (*types.Message, error) {
	// Run the Query
	if err := r.db.Create(message).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Message
	return message, nil
}
