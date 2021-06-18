package server

import (
	"context"
	"gochat/platform/internal/proto/pb"
)

// Login gRPC Method
func (s *AccountService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Call the Controller Method
	loginResponse, err := s.c.Login(in)
	if err != nil {
		return nil, err
	}

	// Return the LoginResponse
	return loginResponse, nil
}
