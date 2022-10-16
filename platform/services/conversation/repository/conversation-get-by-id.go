package repository

import (
	"fmt"
	"gochat/platform/services/conversation/types"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// GetConversationByID fetches a conversation by its UUID
func (r *ConversationRepository) GetConversationByID(conversationID uuid.UUID) (*types.Conversation, error) {
	// Create the Return Conversation
	var rtnConversation *types.Conversation

	// Run the Query
	if err := r.db.Model(&types.Conversation{}).Preload(clause.Associations).First(&rtnConversation, "id = ?", conversationID).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation
	return rtnConversation, nil
}
