package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/conversation/repository/types"
)

// UpdateConversation updates a conversation object
func (c *ConversationController) UpdateConversation(in *pb.Conversation) (*pb.Conversation, error) {
	if in.Id == "" {
		return nil, fmt.Errorf("controller error: no ID specified in UpdateConversation")
	}

	// Convert to a Type
	conversation := &types.Conversation{}
	if err := pbjson.FromProto(in, conversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Check that the Conversation Exists
	_, err := c.r.GetConversationByID(conversation.ID)
	if err != nil {
		return nil, err
	}

	// Update the Conversation
	updatedConversation, err := c.r.UpdateConversation(conversation)
	if err != nil {
		return nil, err
	}

	// Convert back to Proto
	rtnConversation := &pb.Conversation{}
	if err := pbjson.ToProto(updatedConversation, rtnConversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the Updated Conversation
	return rtnConversation, nil
}
