package types

import (
	"time"

	"github.com/gofrs/uuid"
)

// Message database model
type Message struct {
	ID             uuid.UUID `gorm:"type:text" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	ConversationID uuid.UUID `gorm:"foreign_key:conversation_id" json:"conversation_id"`
	AuthorID       uuid.UUID `gorm:"author_id;type:text" json:"author_id"`
	Message        string    `json:"message"`
}
