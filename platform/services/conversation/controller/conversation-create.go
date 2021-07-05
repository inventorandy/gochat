package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/conversation/repository/types"
	"regexp"
	"strings"

	"github.com/gofrs/uuid"
)

// CreateConversation creates a new conversation object
func (c *ConversationController) CreateConversation(in *pb.Conversation) (*pb.Conversation, error) {
	// Check whether the conversation is public or not
	if in.IsPublic {
		// Check that we have a label applied for a public conversation
		if in.Label == "" {
			return nil, fmt.Errorf("controller error: no conversation label provided in CreateConversation")
		}

		// Remove non alpha-numeric characters and make lowercase
		re := regexp.MustCompile("([[:^alnum:]])")
		in.Label = re.ReplaceAllString(in.Label, "")
		in.Label = strings.ToLower(in.Label)

		// Preceed the Label with a hashtag if not already done
		/*
			if !strings.HasPrefix(in.Label, "#") {
				in.Label = fmt.Sprintf("#%s", in.Label)
			}
		*/
	}

	// Create a new ID
	conversationID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("controller error: unable to create new UUID in CreateConversation")
	}

	// Convert the Conversation Proto to Gorm Type
	conversation := &types.Conversation{}
	if err := pbjson.FromProto(in, conversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Set the ID
	conversation.ID = conversationID

	// Run the Repository Method
	newConversation, err := c.r.CreateConversation(conversation)
	if err != nil {
		return nil, err
	}

	// Convert back to Proto
	rtnConversation := &pb.Conversation{}
	if err := pbjson.ToProto(newConversation, rtnConversation); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the Conversation
	return rtnConversation, nil
}
