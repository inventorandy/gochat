package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// GetConversationByID gRPC Method
func (s *ConversationService) GetConversationByID(ctx context.Context, in *wrappers.StringValue) (*pb.Conversation, error) {
	// Call the Controller Method
	conversation, err := s.c.GetConversationByID(in)
	if err != nil {
		return nil, err
	}

	// Return the Conversation
	return conversation, nil
}
