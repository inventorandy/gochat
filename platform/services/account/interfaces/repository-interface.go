package interfaces

import (
	"gochat/platform/services/account/types"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Repository interface
type Repository interface {
	CreateUser(user *types.User) (*types.User, error)
	GetUserByID(userID uuid.UUID) (*types.User, error)
	GetUserByEmail(email strfmt.Email) (*types.User, error)
	GetAllUsers() ([]*types.User, error)
}
