package controller

import (
	"context"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (c *HandlerController) ConversationGet(user *pb.User, in stable.GetConversationParams) middleware.Responder {
	if *in.Public {
		// Get the Public Conversations
		conversationList, err := c.conversations.GetPublicConversations(context.Background(), &empty.Empty{})
		if err != nil {
			log.Println(err.Error())
			return stable.NewGetConversationBadRequest().WithPayload(&models.Error{
				Message: "Could not find public conversations.",
			})
		}

		// Convert to a Swagger Model
		var rtnConversations models.ConversationList
		for _, conversation := range conversationList.Conversations {
			swaggerConversation := &models.Conversation{}
			if err := pbjson.FromProto(conversation, swaggerConversation); err != nil {
				log.Println(err.Error())
				return stable.NewGetConversationBadRequest().WithPayload(&models.Error{
					Message: "Could not convert public conversations to a readable format.",
				})
			}

			// Add ot the Return List
			rtnConversations = append(rtnConversations, swaggerConversation)
		}

		// Return the List
		return stable.NewGetConversationOK().WithPayload(rtnConversations)
	}

	// Get the Private Conversations
	conversationList, err := c.conversations.GetPrivateConversationsForUser(context.Background(), &wrappers.StringValue{Value: user.Id})
	if err != nil {
		log.Println(err.Error())
		return stable.NewGetConversationBadRequest().WithPayload(&models.Error{
			Message: "Could not find private conversations.",
		})
	}

	// Convert to a Swagger Model
	var rtnConversations models.ConversationList
	for _, conversation := range conversationList.Conversations {
		swaggerConversation := &models.Conversation{}
		if err := pbjson.FromProto(conversation, swaggerConversation); err != nil {
			log.Println(err.Error())
			return stable.NewGetConversationBadRequest().WithPayload(&models.Error{
				Message: "Could not convert private conversations to a readable format.",
			})
		}

		// Add ot the Return List
		rtnConversations = append(rtnConversations, swaggerConversation)
	}

	// Return the List
	return stable.NewGetConversationOK().WithPayload(rtnConversations)
}
