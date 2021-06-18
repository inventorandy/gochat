package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// CreateConversation gRPC Method
func (s *ConversationService) CreateConversation(ctx context.Context, in *pb.Conversation) (*pb.Conversation, error) {
	// Call the Controller Method
	conversation, err := s.c.CreateConversation(in)
	if err != nil {
		return nil, err
	}

	// Send an Event to the ChatStream
	s.sendChatEvent(&pb.ChatStreamUpdate{
		UpdateType: &pb.ChatStreamUpdate_ConversationInfo{
			ConversationInfo: &pb.ConversationUpdate{
				UpdateType:   pb.ConversationUpdateType_CONVERSATION_CREATED,
				Conversation: conversation,
			},
		},
	})

	// Return the Conversation
	return conversation, nil
}
