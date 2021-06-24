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

// AccountPost [POST /user]
func (c *HandlerController) AccountPost(in stable.PostAccountParams) middleware.Responder {
	// Convert the Account Object to Proto
	userProto := &pb.User{}
	if err := pbjson.ToProto(in.User, userProto); err != nil {
		log.Println(err.Error())
		return stable.NewPostAccountBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert account information.",
		})
	}

	// Send the Request
	newUser, err := c.accounts.CreateUser(context.Background(), userProto)
	if err != nil {
		log.Println(err.Error())
		return stable.NewPostAccountBadRequest().WithPayload(&models.Error{
			Message: "Unable to save account information.",
		})
	}

	// Convert the User back to Swagger
	rtnUser := &models.User{}
	if err := pbjson.FromProto(newUser, rtnUser); err != nil {
		log.Println(err.Error())
		return stable.NewPostAccountInternalServerError().WithPayload(&models.Error{
			Message: "Unable to convert account information.",
		})
	}

	// Return the User Account
	return stable.NewPostAccountCreated().WithPayload(rtnUser)
}
