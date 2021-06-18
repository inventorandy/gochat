package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// GetPrivateConversationsForUser gRPC Method
func (s *ConversationService) GetPrivateConversationsForUser(ctx context.Context, in *wrappers.StringValue) (*pb.ConversationList, error) {
	// Call the Controller Method
	conversations, err := s.c.GetPrivateConversationsForUser(in)
	if err != nil {
		return nil, err
	}

	// Return the Conversations
	return conversations, nil
}
