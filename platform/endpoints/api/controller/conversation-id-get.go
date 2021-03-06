package controller

import (
	"context"
	"fmt"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// ConversationIDGet [GET /conversation/{id}]
func (c *HandlerController) ConversationIDGet(user *pb.User, in stable.GetConversationIDParams) middleware.Responder {
	// Get the Conversation
	conversation, err := c.conversations.GetConversationByID(context.Background(), &wrappers.StringValue{Value: in.ID})
	if err != nil {
		log.Println(err.Error())
		return stable.NewGetConversationIDNotFound().WithPayload(&models.Error{
			Message: "Could not find this conversation in the system.",
		})
	}

	// Convert to swagger
	rtnConversation := &models.Conversation{}
	if err := pbjson.FromProto(conversation, rtnConversation); err != nil {
		log.Println(err.Error())
		return stable.NewGetConversationIDBadRequest().WithPayload(&models.Error{
			Message: "Could not convert conversation to a readable format.",
		})
	}

	// Loop through the Conversation Messages
	if len(rtnConversation.Messages) > 0 {
		for i, message := range rtnConversation.Messages {
			// Get the User
			author, err := c.accounts.GetUserByID(context.Background(), &wrappers.StringValue{Value: message.AuthorID.String()})
			if err != nil {
				log.Println(err.Error())
				rtnConversation.Messages[i].AuthorName = "Unknown Sender"
			} else {
				// Set the Author
				rtnConversation.Messages[i].AuthorName = fmt.Sprintf("%s %s", author.FirstName, author.LastName)
			}
		}
	}

	// Check if the conversation is public
	if *rtnConversation.IsPublic {
		// Return the Conversation
		return stable.NewGetConversationIDOK().WithPayload(rtnConversation)
	}

	// Loop through and see if user is on the participant list
	for _, participant := range conversation.Participants {
		// Check for a match
		if participant.ParticipantId == user.Id {
			// Return the Conversation
			return stable.NewGetConversationIDOK().WithPayload(rtnConversation)
		}
	}

	// Return Unauthorised by Default
	return stable.NewGetConversationIDUnauthorized().WithPayload(&models.Error{
		Message: "You are not authorised to view this conversation.",
	})
}
