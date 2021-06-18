package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// GetUserByEmail gRPC Method
func (s *AccountService) GetUserByEmail(ctx context.Context, in *wrappers.StringValue) (*pb.User, error) {
	// Call the Controller Method
	user, err := s.c.GetUserByEmail(in)
	if err != nil {
		return nil, err
	}

	// Return the User
	return user, nil
}
