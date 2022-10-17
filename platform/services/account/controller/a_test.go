package controller

import (
	"fmt"
	"gochat/platform/services/account/interfaces/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// Suite is a global struct shared by all tests
type Suite struct {
	suite.Suite
	Repository *mocks.MockRepository
	Controller *AccountController
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

// SetupSuite initialises the test suite with a mock repository
func (s *Suite) SetupSuite() {
	// Setup the mock repository
	s.Repository = mocks.NewMockRepository(gomock.NewController(s.T()))

	// Create the controller
	s.Controller = &AccountController{
		r: s.Repository,
	}
}

// TearDownSuite does nothing currently but can be used to close database connections
func (s *Suite) TearDownSuite() {
	fmt.Println("Tearing down suite")
}
