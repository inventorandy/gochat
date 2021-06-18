package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// GetUserByID gRPC Method
func (s *AccountService) GetUserByID(ctx context.Context, in *wrappers.StringValue) (*pb.User, error) {
	// Call the Controller Method
	user, err := s.c.GetUserByID(in)
	if err != nil {
		return nil, err
	}

	// Return the User
	return user, nil
}
