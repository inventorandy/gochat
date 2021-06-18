package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// CreateMessage gRPC Method
func (s *ConversationService) CreateMessage(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	// Call the Controller Method
	message, err := s.c.CreateMessage(in)
	if err != nil {
		return nil, err
	}

	// Send an Event to the ChatStream
	s.sendChatEvent(&pb.ChatStreamUpdate{
		UpdateType: &pb.ChatStreamUpdate_MessageInfo{
			MessageInfo: &pb.MessageUpdate{
				UpdateType: pb.MessageUpdateType_MESSAGE_CREATED,
				Message:    message,
			},
		},
	})

	// Return the Message
	return message, nil
}
