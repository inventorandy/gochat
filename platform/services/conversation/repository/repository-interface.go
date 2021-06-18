package repository

import (
	"gochat/platform/services/conversation/repository/types"

	"github.com/gofrs/uuid"
)

// Repository interface
type Repository interface {
	// Conversation Methods
	CreateConversation(conversation *types.Conversation) (*types.Conversation, error)
	GetConversationByID(conversationID uuid.UUID) (*types.Conversation, error)
	GetPublicConversations() ([]*types.Conversation, error)
	GetPrivateConversationsForUser(userID uuid.UUID) ([]*types.Conversation, error)
	AddUserToConversation(userID uuid.UUID, conversationID uuid.UUID) error

	// Message Methods
	CreateMessage(message *types.Message) (*types.Message, error)
}
