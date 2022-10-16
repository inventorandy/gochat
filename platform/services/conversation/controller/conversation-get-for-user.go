package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

// GetPrivateConversationsForUser method fetches a list of private conversations for a user
func (c *ConversationController) GetPrivateConversationsForUser(in *wrappers.StringValue) (*pb.ConversationList, error) {
	// Check that we have a User ID specified
	if in.Value == "" {
		return nil, fmt.Errorf("controller error: no user Id specified in GetPrivateConversationsForUser")
	}

	// Convert to a UUID
	userID, err := uuid.Parse(in.Value)
	if err != nil {
		return nil, fmt.Errorf("controller error: invalid user Id specified in GetPrivateConversationsForUser")
	}

	// Call the Repository Method
	conversations, err := c.r.GetPrivateConversationsForUser(userID)
	if err != nil {
		return nil, err
	}

	// Create the Return List
	rtn := &pb.ConversationList{}

	// Loop through and add the Conversations
	for _, conversation := range conversations {
		// Convert to Proto
		protoConversation := &pb.Conversation{}
		if err := pbjson.ToProto(conversation, protoConversation); err != nil {
			return nil, fmt.Errorf("controller error: %s", err.Error())
		}

		// Add to the List
		rtn.Conversations = append(rtn.Conversations, protoConversation)
	}

	// Return the List
	return rtn, nil
}
