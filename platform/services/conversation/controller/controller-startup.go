package controller

import (
	"gochat/platform/services/conversation/repository"
)

// ConversationController struct
type ConversationController struct {
	r repository.Repository
}

// NewConversationController creates a new instance of the conversation controller
func NewConversationController() (*ConversationController, error) {
	// Create an instance of the Repository
	repo, err := repository.NewConversationRepository()
	if err != nil {
		return nil, err
	}

	// Create the Controller Object
	cont := &ConversationController{
		r: repo,
	}

	// Return the Controller
	return cont, nil
}
