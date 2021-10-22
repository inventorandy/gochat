package controller

import (
	"gochat/platform/services/account/repository"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

// Claims struct for JWT Claims
type Claims struct {
	UserID uuid.UUID
	jwt.StandardClaims
}

// AccountController struct
type AccountController struct {
	r repository.Repository
}

// NewAccountController creates a new instance of the account controller
func NewAccountController() (*AccountController, error) {
	// Create an instance of the Repository
	repo, err := repository.NewAccountRepository()
	if err != nil {
		return nil, err
	}

	// Create the Controller Object
	cont := &AccountController{
		r: repo,
	}

	// Return the Controller
	return cont, nil
}
