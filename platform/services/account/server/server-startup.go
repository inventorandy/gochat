package server

import (
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/account/controller"
)

// AccountService struct
type AccountService struct {
	c *controller.AccountController
	pb.AccountServiceServer
}

// NewAccountService creates a new instance of the Account gRPC Server
func NewAccountService() (*AccountService, error) {
	// Start up the Controller
	controller, err := controller.NewAccountController()
	if err != nil {
		return nil, err
	}

	// Return the new gRPC Server
	return &AccountService{
		c: controller,
	}, nil
}
