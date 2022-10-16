package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

// GetConversationByID fetches a conversation by its ID
func (c *ConversationController) GetConversationByID(in *wrappers.StringValue) (*pb.Conversation, error) {
	// Check that we have an ID posted
	if in.Value == "" {
		return nil, fmt.Errorf("controller error: no ID specified in GetConversationByID")
	}

	// Convert to a UUID
	conversationID, err := uuid.Parse(in.Value)
	if err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Run the Query
	conversation, err := c.r.GetConversationByID(conversationID)
	if err != nil {
		return nil, err
	}

	// Convert to a Proto
	rtnConversation := &pb.Conversation{}
	if err := pbjson.ToProto(conversation, rtnConversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the Conversation
	return rtnConversation, nil
}
