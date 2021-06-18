package repository

import (
	"fmt"
	"gochat/platform/services/account/repository/types"

	"github.com/go-openapi/strfmt"
)

// GetUserByEmail fetches a user by their email address
func (r *AccountRepository) GetUserByEmail(email strfmt.Email) (*types.User, error) {
	// Create the Return Object
	var user *types.User

	// Run the Query
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Result
	return user, nil
}
