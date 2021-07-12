package types

import (
	"time"

	"github.com/gofrs/uuid"
)

// Conversation database model
type Conversation struct {
	ID            uuid.UUID `gorm:"type:text" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	LastMessageOn time.Time                     `json:"last_message_on"`
	Label         string                        `gorm:"type:text" json:"label"`
	IsPublic      *bool                         `gorm:"type:boolean" json:"is_public"`
	Participants  []*ConversationHasParticipant `gorm:"foreign_key:conversation_id" json:"participants"`
	Messages      []*Message                    `gorm:"foreign_key:conversation_id" json:"messages"`
}

// ConversationHasParticipant database model
type ConversationHasParticipant struct {
	ConversationID uuid.UUID     `gorm:"type:text" json:"conversation_id"`
	ParticipantID  uuid.UUID     `gorm:"type:text" json:"participant_id"`
	Conversation   *Conversation `gorm:"foreign_key:conversation_id"`
}
