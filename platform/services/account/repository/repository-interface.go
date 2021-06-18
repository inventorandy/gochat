package repository

import (
	"gochat/platform/services/account/repository/types"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
)

// Repository interface
type Repository interface {
	CreateUser(user *types.User) (*types.User, error)
	GetUserByID(userID uuid.UUID) (*types.User, error)
	GetUserByEmail(email strfmt.Email) (*types.User, error)
}
