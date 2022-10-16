package controller

import (
	"fmt"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

// UserHasAccessToConversation checks whether a user can read or write messages in a conversation
func (c *ConversationController) UserHasAccessToConversation(in *pb.UserAccessQuery) (*wrappers.BoolValue, error) {
	// Check that we have a User ID posted
	if in.UserId == "" {
		return nil, fmt.Errorf("controller error: no user ID specified in UserHasAccessToConversation")
	}

	// Convert to a UUID
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Check that we have a Conversation ID posted
	if in.ConversationId == "" {
		return nil, fmt.Errorf("controller error: no conversation ID speciied in UserHasAccessToConversation")
	}

	// Convert to a UUID
	conversationID, err := uuid.Parse(in.ConversationId)
	if err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Get the Conversation
	conversation, err := c.r.GetConversationByID(conversationID)
	if err != nil {
		return nil, err
	}

	// Check if the Conversation is Public
	if *conversation.IsPublic {
		// If it is, the user has access
		return &wrappers.BoolValue{Value: true}, nil
	}

	// Check if the user is a participant in the conversation
	_, err = c.r.GetConversationParticipantByID(conversationID, userID)
	if err != nil {
		// We got a NotFound (or repo error), return FALSE
		return &wrappers.BoolValue{Value: false}, nil
	}

	// We made it this far, the user has access
	return &wrappers.BoolValue{Value: true}, nil
}
