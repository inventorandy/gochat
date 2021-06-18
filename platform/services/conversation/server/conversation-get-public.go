package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

// GetPublicConversations gRPC Method
func (s *ConversationService) GetPublicConversations(ctx context.Context, in *empty.Empty) (*pb.ConversationList, error) {
	// Call the Controller Method
	conversations, err := s.c.GetPublicConversations(in)
	if err != nil {
		return nil, err
	}

	// Return the Conversations
	return conversations, nil
}
