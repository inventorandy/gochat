package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"

	"github.com/gofrs/uuid"
)

// GetConversationParticipantByID fetches a conversation participant by their User ID
func (r *ConversationRepository) GetConversationParticipantByID(conversationID uuid.UUID, userID uuid.UUID) (*types.ConversationHasParticipant, error) {
	// Create the Return Object
	var conversationParticipant *types.ConversationHasParticipant

	// Run the Query
	if err := r.db.First(&conversationParticipant, "conversation_id = ? AND participant_id = ?", conversationID, userID).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Result
	return conversationParticipant, nil
}
