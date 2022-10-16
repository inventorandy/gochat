package types

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// User database model
type User struct {
	ID        uuid.UUID `gorm:"type:text" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     strfmt.Email `gorm:"type:text" json:"email"`
	Password  string       `json:"password"`
	Salt      string       `json:"salt"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
}
