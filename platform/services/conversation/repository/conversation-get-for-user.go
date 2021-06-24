package repository

import (
	"fmt"
	"gochat/platform/services/conversation/repository/types"

	"github.com/gofrs/uuid"
	"gorm.io/gorm/clause"
)

// GetPrivateConversationsForUser fetches a list of private conversations for a user
func (r *ConversationRepository) GetPrivateConversationsForUser(userID uuid.UUID) ([]*types.Conversation, error) {
	// Create the Return Object
	var rtnConversations []*types.Conversation

	// Run the Query
	if err := r.db.Model(&types.ConversationHasParticipant{}).
		Where("conversation_has_participants.participant_id = ?", userID).
		Association(clause.Associations).
		Find(rtnConversations); err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Conversation List
	return rtnConversations, nil
}
