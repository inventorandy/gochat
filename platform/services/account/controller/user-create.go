package controller

import (
	"fmt"
	"gochat/platform/internal/enc"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/account/types"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// CreateUser validates a user proto and creates a user in the database via the repository
func (c *AccountController) CreateUser(in *pb.User) (*pb.User, error) {
	// Check that we have an email posted
	if in.Email == "" {
		return nil, fmt.Errorf("controller error: no email specified in CreateUser")
	}

	// Check that the email is valid
	if !strfmt.IsEmail(in.Email) {
		return nil, fmt.Errorf("controller error: invalid email specified in CreateUser")
	}

	// Check for an existing user
	existingUser, _ := c.r.GetUserByEmail(strfmt.Email(in.Email))
	if existingUser != nil {
		return nil, fmt.Errorf("controller error: user with this email already exists in CreateUser")
	}

	// Check that we have a password posted
	if in.Password == "" {
		return nil, fmt.Errorf("controller error: no password specified in CreateUser")
	}

	// Check that we have a first name posted
	if in.FirstName == "" {
		return nil, fmt.Errorf("controller error: no first name specified in CreateUser")
	}

	// Generate a new ID
	userID := uuid.New()

	// Convert the User to a Type
	userIn := &types.User{}
	if err := pbjson.FromProto(in, userIn); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Add the New ID
	userIn.ID = userID

	// Generate a Salt for the Password (32 bytes / 256 bits)
	userIn.Salt = enc.GenerateSalt(32)

	// Encrypt the Password with the Salt using SHA256
	userIn.Password = enc.SHA256EncryptWithSalt(in.Password, userIn.Salt)

	// Call the Repository Method
	user, err := c.r.CreateUser(userIn)
	if err != nil {
		return nil, err
	}

	// Convert back to a Proto
	userOut := &pb.User{}
	if err := pbjson.ToProto(user, userOut); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the User
	return userOut, nil
}
