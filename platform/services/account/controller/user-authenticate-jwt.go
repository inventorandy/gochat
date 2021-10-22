package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// AuthenticateJWT authenticates a login JWT and returns a user when successful
func (c *AccountController) AuthenticateJWT(in *wrappers.StringValue) (*pb.User, error) {
	// Check that we have a JWT
	if in.Value == "" {
		return nil, fmt.Errorf("controller error: no JWT specified in AuthenticateJWT")
	}

	// Get the Signing Key
	jwtKey := []byte(os.Getenv("JWT_SIGNATURE"))

	// Create the Claims
	claims := &Claims{}

	// Parse the Token
	tkn, err := jwt.ParseWithClaims(in.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for Errors
	if err != nil {
		// Check for Invalid Signature
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid JWT signature in token")
		}

		// Return generic error otherwise
		return nil, fmt.Errorf("error parsing JWT: %s", err.Error())
	}

	// Check for Validity of Token
	if !tkn.Valid {
		// Return an Error
		return nil, fmt.Errorf("JWT is no longer valid")
	}

	// Get the User
	user, err := c.r.GetUserByID(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("controller error: user not found")
	}

	// Convert to a Proto
	userOut := &pb.User{}
	if err := pbjson.ToProto(user, userOut); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Return the User
	return userOut, nil
}
