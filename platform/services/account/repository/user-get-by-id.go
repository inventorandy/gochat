package repository

import (
	"fmt"
	"gochat/platform/services/account/repository/types"

	"github.com/gofrs/uuid"
)

// GetUserByID fetches a user by their UUID
func (r *AccountRepository) GetUserByID(userID uuid.UUID) (*types.User, error) {
	// Create the Return Object
	var user *types.User

	// Run the Query
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Result
	return user, nil
}
