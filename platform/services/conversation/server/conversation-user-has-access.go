package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// UserHasAccessToConversation gRPC Method
func (s *ConversationService) UserHasAccessToConversation(ctx context.Context, in *pb.UserAccessQuery) (*wrappers.BoolValue, error) {
	// Call the Controller Method
	userHasAccess, err := s.c.UserHasAccessToConversation(in)
	if err != nil {
		return nil, err
	}

	// Return the Result
	return userHasAccess, nil
}
