package controller

import (
	"context"
	"fmt"
	"gochat/platform/internal/proto/pb"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// AuthenticateJWT authenticates a JWT and returns the user
func (c *HandlerController) AuthenticateJWT(jwt string) (*pb.User, error) {
	// Check if a JWT was posted
	if jwt == "" {
		return nil, fmt.Errorf("no Authorization header set")
	}

	// Call the Service Method
	user, err := c.accounts.AuthenticateJWT(context.Background(), &wrapperspb.StringValue{Value: jwt})
	if err != nil {
		return nil, err
	}

	// Return the User
	return user, nil
}
