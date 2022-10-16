package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

// GetUserByID fetches a user by their ID
func (c *AccountController) GetUserByID(in *wrappers.StringValue) (*pb.User, error) {
	// Check that a UUID has been posted
	if in.Value == "" {
		return nil, fmt.Errorf("controller error: no ID specified in GetUserByID")
	}

	// Attempt to convert to a UUID
	userID, err := uuid.Parse(in.Value)
	if err != nil {
		return nil, fmt.Errorf("controller error: unable to convert UUID from string: %s", err.Error())
	}

	// Call the Repository Method
	user, err := c.r.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Convert to a Proto
	userOut := &pb.User{}
	if err := pbjson.ToProto(user, userOut); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Remove the Password and Salt from the Return
	userOut.Password = ""
	userOut.Salt = ""

	// Return the User
	return userOut, nil
}
