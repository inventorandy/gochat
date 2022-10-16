package repository

import (
	"fmt"
	"gochat/platform/services/account/types"
)

// GetAllUsers fetches a list of all users
func (r *AccountRepository) GetAllUsers() ([]*types.User, error) {
	// Create the Return Object
	var rtnUsers []*types.User

	// Run the Query
	if err := r.db.Order("first_name ASC").Find(&rtnUsers).Error; err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	// Return the Results
	return rtnUsers, nil
}
