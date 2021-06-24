package controller

import (
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

// AccountGet [GET /account]
func (c *HandlerController) AccountGet(user *pb.User, in stable.GetAccountParams) middleware.Responder {
	// Convert the User to a Swagger
	rtnUser := &models.User{}
	if err := pbjson.FromProto(user, rtnUser); err != nil {
		log.Println(err.Error())
		return stable.NewGetAccountInternalServerError().WithPayload(&models.Error{
			Message: "Could not convert user to readable format.",
		})
	}

	// Return the User
	return stable.NewGetAccountOK().WithPayload(rtnUser)
}
