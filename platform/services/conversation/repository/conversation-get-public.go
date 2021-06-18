package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"

	"gorm.io/gorm/clause"
)

// GetPublicConversations gets a list of public conversations on the platform
func (r *ConversationRepository) GetPublicConversations() ([]*types.Conversation, error) {
	// Create the Return Array
	var rtnConversations []*types.Conversation

	// Run the Query
	if err := r.db.Preload(clause.Associations).Order("label ASC").Find(&rtnConversations, "is_public", true).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversations
	return rtnConversations, nil
}
