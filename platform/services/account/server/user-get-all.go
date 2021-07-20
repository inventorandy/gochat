package server

import (
	"context"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

// GetAllUsers gRPC Method
func (s *AccountService) GetAllUsers(ctx context.Context, in *empty.Empty) (*pb.UserList, error) {
	// Call the Controller Method
	users, err := s.c.GetAllUsers(in)
	if err != nil {
		return nil, err
	}

	// Return the User List
	return users, nil
}
