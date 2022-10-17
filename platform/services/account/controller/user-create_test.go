package controller

import (
	"fmt"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/account/types"

	"github.com/golang/mock/gomock"
)

// TestCreateUser_OK tests a good implementation
func (s *Suite) TestCreateUser_OK() {
	// Build the Input Object
	userPB := &pb.User{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "Testy",
		LastName:  "McTestface",
	}
	userDB := &types.User{
		Email:     "test@test.com",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	// Mock the Repository Calls
	gomock.InOrder(
		s.Repository.EXPECT().GetUserByEmail(userDB.Email).Times(1).Return(nil, fmt.Errorf("repository error")),
		s.Repository.EXPECT().CreateUser(gomock.Any()).Times(1).Return(userDB, nil),
	)

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that no error was returned
	s.Assert().NoError(err)

	// Assert that the User was created
	s.Assert().NotNil(user)
	s.Assert().Equal(userPB.FirstName, user.FirstName)
	s.Assert().Equal(userPB.LastName, user.LastName)
	s.Assert().Equal(userPB.Email, user.Email)
}

// TestCreateUser_NoEmail tests no email
func (s *Suite) TestCreateUser_NoEmail() {
	// Build the Input Object
	userPB := &pb.User{
		Password:  "password123",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that an error was returned
	s.Assert().Error(err)

	// Assert that no user was returned
	s.Assert().Nil(user)
}

// TestCreateUser_InvalidEmail tests an invalid email
func (s *Suite) TestCreateUser_InvalidEmail() {
	// Build the Input Object
	userPB := &pb.User{
		Email:     "notanemailaddress",
		Password:  "password123",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that an error was returned
	s.Assert().Error(err)

	// Assert that no user was returned
	s.Assert().Nil(user)
}

// TestCreateUser_UserExists tests a user already existing
func (s *Suite) TestCreateUser_UserExists() {
	// Build the Input Object
	userPB := &pb.User{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "Testy",
		LastName:  "McTestface",
	}
	userDB := &types.User{
		Email:     "test@test.com",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	// Mock the Repository Calls
	gomock.InOrder(
		s.Repository.EXPECT().GetUserByEmail(userDB.Email).Times(1).Return(userDB, nil),
	)

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that an error was returned
	s.Assert().Error(err)

	// Assert that no user was returned
	s.Assert().Nil(user)
}

// TestCreateUser_NoPassword tests no password
func (s *Suite) TestCreateUser_NoPassword() {
	// Build the Input Object
	userPB := &pb.User{
		Email:     "test@test.com",
		FirstName: "Testy",
		LastName:  "McTestface",
	}
	userDB := &types.User{
		Email:     "test@test.com",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	// Mock the Repository Calls
	gomock.InOrder(
		s.Repository.EXPECT().GetUserByEmail(userDB.Email).Times(1).Return(nil, fmt.Errorf("repository error")),
	)

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that an error was returned
	s.Assert().Error(err)

	// Assert that no user was returned
	s.Assert().Nil(user)
}

// TestCreateUser_NoFirstName tests no first name
func (s *Suite) TestCreateUser_NoFirstName() {
	// Build the Input Object
	userPB := &pb.User{
		Email:    "test@test.com",
		Password: "password123",
		LastName: "McTestface",
	}
	userDB := &types.User{
		Email:    "test@test.com",
		Password: "password123",
		LastName: "McTestface",
	}

	// Mock the Repository Calls
	gomock.InOrder(
		s.Repository.EXPECT().GetUserByEmail(userDB.Email).Times(1).Return(nil, fmt.Errorf("repository error")),
	)

	// Call the Controller Method
	user, err := s.Controller.CreateUser(userPB)

	// Assert that an error was returned
	s.Assert().Error(err)

	// Assert that no user was returned
	s.Assert().Nil(user)
}
