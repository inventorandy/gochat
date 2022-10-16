package repository

import (
	"fmt"
	"gochat/platform/services/conversation/types"

	"github.com/google/uuid"
)

// AddUserToConversation adds a user to a conversation
func (r *ConversationRepository) AddUserToConversation(userID uuid.UUID, conversationID uuid.UUID) error {
	// Create the DB Object
	assocObj := &types.ConversationHasParticipant{
		ConversationID: conversationID,
		ParticipantID:  userID,
	}

	// Run the Query
	if err := r.db.Create(assocObj).Error; err != nil {
		return fmt.Errorf("repository error: %s", err.Error())
	}

	// Return Nil by Default
	return nil
}
