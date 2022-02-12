package controller

import (
	"fmt"
	"gochat/platform/internal/enc"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt"
)

// Login checks a user's credentials and returns a JWT on success
func (c *AccountController) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Check that an Email was posted
	if in.Email == "" {
		return nil, fmt.Errorf("no email address specified")
	}

	// Check that the Email was valid
	if !strfmt.IsEmail(in.Email) {
		return nil, fmt.Errorf("email address was invalid")
	}

	// Check that a Password was posted
	if in.Password == "" {
		return nil, fmt.Errorf("no password was specified")
	}

	// Get the User
	user, err := c.r.GetUserByEmail(strfmt.Email(in.Email))
	if err != nil {
		return nil, fmt.Errorf("user with this email could not be found")
	}

	// Hash the Input Password
	inPass := enc.SHA256EncryptWithSalt(in.Password, user.Salt)

	// Compare the Passwords
	if inPass != user.Password {
		return nil, fmt.Errorf("the password entered was incorrect")
	}

	// Convert to a Proto
	userOut := &pb.User{}
	if err := pbjson.ToProto(user, userOut); err != nil {
		return nil, fmt.Errorf("controller error: %s", err.Error())
	}

	// Remove the Password and Salt from the Return
	userOut.Password = ""
	userOut.Salt = ""

	// Set the JWT Expiration Time
	expirationTime := time.Now().Add(2 * time.Hour)

	// Create the JWT Claims
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the Token with the Algorithm for Signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))
	if err != nil {
		return nil, fmt.Errorf("unable to generate a login token")
	}

	// Otherwise, create the return object
	loginResponse := &pb.LoginResponse{
		AuthToken: tokenStr,
		User:      userOut,
	}

	// Return the Login Response
	return loginResponse, nil
}
