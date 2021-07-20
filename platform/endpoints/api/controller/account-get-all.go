package controller

import (
	"context"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/protobuf/ptypes/empty"
)

// AccountGetAll [GET /account/get-all]
func (c *HandlerController) AccountGetAll(in stable.GetAccountGetAllParams) middleware.Responder {
	// Get the User List
	userList, err := c.accounts.GetAllUsers(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println(err.Error())
		return stable.NewGetAccountGetAllInternalServerError().WithPayload(&models.Error{
			Message: "Unable to fetch a list of users on the system.",
		})
	}

	// Now Convert to Swagger
	rtnUsers := models.UserList{}
	for _, userPB := range userList.Users {
		// Convert the User to a Swagger Object and Append
		user := &models.User{}
		if err := pbjson.FromProto(userPB, user); err != nil {
			log.Println(err.Error())
			return stable.NewGetAccountGetAllInternalServerError().WithPayload(&models.Error{
				Message: "Unable to convert user to readable format.",
			})
		}

		// Remove Password and Salt
		user.Password = ""
		user.Salt = ""

		// Append to the List
		rtnUsers = append(rtnUsers, user)
	}

	// Return the List
	return stable.NewGetAccountGetAllOK().WithPayload(rtnUsers)
}
