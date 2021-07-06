package controller

import (
	"context"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// MessagePost [POST /message]
func (c *HandlerController) MessagePost(user *pb.User, in stable.PostMessageParams) middleware.Responder {
	// Set the Message Author
	in.Message.AuthorID = strfmt.UUID(user.Id)

	// Convert the Message Object to Proto
	messageProto := &pb.Message{}
	if err := pbjson.ToProto(in.Message, messageProto); err != nil {
		log.Println(err.Error())
		return stable.NewPostMessageBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert message information.",
		})
	}

	// Send the Request
	newMessage, err := c.conversations.CreateMessage(context.Background(), messageProto)
	if err != nil {
		log.Println(err.Error())
		return stable.NewPostMessageBadRequest().WithPayload(&models.Error{
			Message: "Unable to save message information.",
		})
	}

	// Convert the Message back to Swagger
	rtnMessage := &models.Message{}
	if err := pbjson.FromProto(newMessage, rtnMessage); err != nil {
		log.Println(err.Error())
		return stable.NewPostMessageBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert message information.",
		})
	}

	// Return the Message
	return stable.NewPostMessageCreated().WithPayload(rtnMessage)
}
