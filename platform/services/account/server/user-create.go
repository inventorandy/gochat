package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// CreateUser gRPC Method
func (s *AccountService) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	// Call the Controller Method
	user, err := s.c.CreateUser(in)
	if err != nil {
		return nil, err
	}

	// Return the User
	return user, nil
}
