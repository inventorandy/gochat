package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

// GetPublicConversations fetches a list of public conversations
func (c *ConversationController) GetPublicConversations(in *empty.Empty) (*pb.ConversationList, error) {
	// Call the Repository Method
	conversations, err := c.r.GetPublicConversations()
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
