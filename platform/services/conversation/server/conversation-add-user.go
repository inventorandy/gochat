package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// AddUserToConversation gRPC Method
func (s *ConversationService) AddUserToConversation(ctx context.Context, in *pb.ConversationHasParticipant) (*pb.Conversation, error) {
	// Call the Controller Method
	conversation, err := s.c.AddUserToConversation(in)
	if err != nil {
		return nil, err
	}

	// Send an Event to the ChatStream
	s.sendChatEvent(&pb.ChatStreamUpdate{
		UpdateType: &pb.ChatStreamUpdate_ConversationInfo{
			ConversationInfo: &pb.ConversationUpdate{
				UpdateType:   pb.ConversationUpdateType_CONVERSATION_UPDATED,
				Conversation: conversation,
			},
		},
	})

	// Return the Conversation
	return conversation, nil
}
