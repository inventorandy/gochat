package repository

import (
	"fmt"
	"gochat/platform/services/account/types"
)

// CreateUser creates a new user in the database
func (r *AccountRepository) CreateUser(user *types.User) (*types.User, error) {
	// Run the Query
	if err := r.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the User
	return user, nil
}
