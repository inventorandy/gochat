package repository

import (
	"fmt"
	"gochat/platform/services/conversation/types"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// GetPrivateConversationsForUser fetches a list of private conversations for a user
func (r *ConversationRepository) GetPrivateConversationsForUser(userID uuid.UUID) ([]*types.Conversation, error) {
	// Conversation ID List
	var conversationsList []*types.ConversationHasParticipant
	if err := r.db.Find(&conversationsList, "participant_id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Only proceed if we have conversations to return
	if len(conversationsList) == 0 {
		return nil, fmt.Errorf("repository error: no private conversations for user %s", userID)
	}

	// Now Create the Array of IDs
	conversationIDs := []uuid.UUID{}
	for _, conversation := range conversationsList {
		conversationIDs = append(conversationIDs, conversation.ConversationID)
	}

	// Create the Return Object
	var rtnConversations []*types.Conversation

	// Run the Query
	if err := r.db.Preload(clause.Associations).Find(&rtnConversations, "is_public != true AND id IN (?)", conversationIDs).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation List
	return rtnConversations, nil
}
