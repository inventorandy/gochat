package repository

import (
	"fmt"
	"gochat/platform/services/conversation/types"

	"gorm.io/gorm/clause"
)

// GetConversationByLabel fetches a conversation by its Label
func (r *ConversationRepository) GetConversationByLabel(conversationLabel string) (*types.Conversation, error) {
	// Create the Return Conversation
	var rtnConversation *types.Conversation

	// Run the Query
	if err := r.db.Model(&types.Conversation{}).Preload(clause.Associations).First(&rtnConversation, "label = ?", conversationLabel).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation
	return rtnConversation, nil
}
