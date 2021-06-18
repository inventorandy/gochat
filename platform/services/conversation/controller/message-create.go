package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/conversation/repository/types"

	"github.com/gofrs/uuid"
)

// CreateMessage creates a new message
func (c *ConversationController) CreateMessage(in *pb.Message) (*pb.Message, error) {
	// Check that the message contains an author ID
	if in.AuthorId == "" {
		return nil, fmt.Errorf("controller error: no author Id specified in CreateMessage")
	}

	// Check that the message contains a conversation ID
	if in.ConversationId == "" {
		return nil, fmt.Errorf("controller error: no conversation Id specified in CreateMessage")
	}

	// Check that the message contains content
	if in.Message == "" {
		return nil, fmt.Errorf("controller error: no message in CreateMessage")
	}

	// Convert to a GORM Type
	message := &types.Message{}
	if err := pbjson.FromProto(in, message); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Check that the Conversation exists
	_, err := c.r.GetConversationByID(message.ConversationID)
	if err != nil {
		return nil, fmt.Errorf("controller error: this conversation does not exist")
	}

	// Create a new ID
	messageID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("controller error: unable to generate new UUID for message: %s", err.Error())
	}

	// Set the ID
	message.ID = messageID

	// Create the Message
	newMessage, err := c.r.CreateMessage(message)
	if err != nil {
		return nil, err
	}

	// Convert back to Proto
	rtnMessage := &pb.Message{}
	if err := pbjson.ToProto(newMessage, rtnMessage); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the Message
	return rtnMessage, nil
}
