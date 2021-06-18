package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/conversation/repository/types"
)

// AddUserToConversation method adds a participant to a conversation
func (c *ConversationController) AddUserToConversation(in *pb.ConversationHasParticipant) (*pb.Conversation, error) {
	// Check that we have a Participant ID specified
	if in.ParticipantId == "" {
		return nil, fmt.Errorf("controller error: no user Id specified in AddUserToConversation")
	}

	// Check that we have a Conversation ID specified
	if in.ConversationId == "" {
		return nil, fmt.Errorf("controller error: no conversation Id specified in AddUserToConversation")
	}

	// Convert to a GORM Type
	rel := &types.ConversationHasParticipant{}
	if err := pbjson.FromProto(in, rel); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Check that the Conversation exists
	_, err := c.r.GetConversationByID(rel.ConversationID)
	if err != nil {
		return nil, err
	}

	// Create the Entry
	if err := c.r.AddUserToConversation(rel.ParticipantID, rel.ConversationID); err != nil {
		return nil, err
	}

	// Get the Conversation
	conversation, err := c.r.GetConversationByID(rel.ConversationID)
	if err != nil {
		return nil, err
	}

	// Convert back to Proto
	rtnConversation := &pb.Conversation{}
	if err := pbjson.ToProto(conversation, rtnConversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the Conversation
	return rtnConversation, nil
}
