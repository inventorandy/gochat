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

// AccountPut [PUT /account]
func (c *HandlerController) AccountPut(user *pb.User, in stable.PutAccountParams) middleware.Responder {
	// Make sure the ID in the input user is the same as the logged in user
	if user.Id != in.User.ID.String() {
		return stable.NewPutAccountUnauthorized().WithPayload(&models.Error{
			Message: "You are not authorised to edit this account.",
		})
	}

	// Conver the Input User to a Proto
	userProto := &pb.User{}
	if err := pbjson.ToProto(in.User, userProto); err != nil {
		log.Println(err.Error())
		return stable.NewPutAccountBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert account information.",
		})
	}

	// Call the Service Method
	updatedUser, err := c.accounts.UpdateUser(context.Background(), userProto)
	if err != nil {
		log.Println(err.Error())
		return stable.NewPutAccountInternalServerError().WithPayload(&models.Error{
			Message: "Unable to save account information.",
		})
	}

	// Convert back to Swagger
	rtnUser := &models.User{}
	if err := pbjson.FromProto(updatedUser, rtnUser); err != nil {
		log.Println(err.Error())
		return stable.NewPutAccountInternalServerError().WithPayload(&models.Error{
			Message: "Unable to convert account information.",
		})
	}

	// Return the User
	return stable.NewPutAccountOK().WithPayload(rtnUser)
}
