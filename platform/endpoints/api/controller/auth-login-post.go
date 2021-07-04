package controller

import (
	"context"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

// AuthLoginPost [POST /auth/login]
func (c *HandlerController) AuthLoginPost(in stable.PostAuthLoginParams) middleware.Responder {
	log.Printf("USER: %v %v\r\n", in.Login.Email, in.Login.Password)
	// Call the Service Method
	loginResponse, err := c.accounts.Login(context.Background(), &pb.LoginRequest{
		Email:    in.Login.Email.String(),
		Password: in.Login.Password,
	})
	if err != nil {
		log.Println(err.Error())
		return stable.NewPostAuthLoginUnauthorized().WithPayload(&models.Error{
			Message: err.Error(),
		})
	}

	// Convert the Login Response to Model
	rtnLogin := &models.LoginResponse{}
	if err := pbjson.FromProto(loginResponse, rtnLogin); err != nil {
		log.Println(err.Error())
		return stable.NewPostAuthLoginInternalServerError().WithPayload(&models.Error{
			Message: err.Error(),
		})
	}

	// Return the Login Response
	return stable.NewPostAuthLoginOK().WithPayload(rtnLogin)
}
