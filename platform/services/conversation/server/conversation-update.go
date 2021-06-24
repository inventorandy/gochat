package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// UpdateConversation gRPC Method
func (s *ConversationService) UpdateConversation(ctx context.Context, in *pb.Conversation) (*pb.Conversation, error) {
	// Call the Controller Method
	conversation, err := s.c.UpdateConversation(in)
	if err != nil {
		return nil, err
	}

	// Return the Conversation
	return conversation, nil
}
