package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/go-openapi/strfmt"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// GetUserByEmail fetches a user by their Email
func (c *AccountController) GetUserByEmail(in *wrappers.StringValue) (*pb.User, error) {
	// Check that a Email has been posted
	if in.Value == "" {
		return nil, fmt.Errorf("controller error: no email specified in GetUserByEmail")
	}

	// Check that the Email is valid
	if !strfmt.IsEmail(in.Value) {
		return nil, fmt.Errorf("controller error: invalid email specified in GetUserByEmail")
	}

	// Attempt to convert to a UUEmail
	userEmail := strfmt.Email(in.Value)

	// Call the Repository Method
	user, err := c.r.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	// Convert to a Proto
	userOut := &pb.User{}
	if err := pbjson.ToProto(user, userOut); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the User
	return userOut, nil
}
