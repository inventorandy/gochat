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

// ConversationPut [PUT /conversation]
func (c *HandlerController) ConversationPut(in stable.PutConversationParams) middleware.Responder {
	// Convert the Conversation Object to Proto
	conversationProto := &pb.Conversation{}
	if err := pbjson.ToProto(in.Conversation, conversationProto); err != nil {
		log.Println(err.Error())
		return stable.NewPutConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert conversation information.",
		})
	}

	// Send the Request
	newConversation, err := c.conversations.UpdateConversation(context.Background(), conversationProto)
	if err != nil {
		log.Println(err.Error())
		return stable.NewPutConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to save conversation information.",
		})
	}

	// Convert the Conversation back to Swagger
	rtnConversation := &models.Conversation{}
	if err := pbjson.FromProto(newConversation, rtnConversation); err != nil {
		log.Println(err.Error())
		return stable.NewPutConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert conversation information.",
		})
	}

	// Return the Conversation
	return stable.NewPutConversationOK().WithPayload(rtnConversation)
}
