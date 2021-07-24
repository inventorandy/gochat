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
)

// ConversationPost [POST /conversation]
func (c *HandlerController) ConversationPost(user *pb.User, in stable.PostConversationParams) middleware.Responder {
	// Clear the Participants
	inConversation := in.Conversation
	inConversation.Participants = nil
	// Convert the Conversation Object to Proto
	conversationProto := &pb.Conversation{}
	if err := pbjson.ToProto(inConversation, conversationProto); err != nil {
		log.Println(err.Error())
		return stable.NewPostConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert conversation information.",
		})
	}

	// Send the Request
	newConversation, err := c.conversations.CreateConversation(context.Background(), conversationProto)
	if err != nil {
		log.Println(err.Error())
		return stable.NewPostConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to save conversation information.",
		})
	}

	// Add the Participants
	if !*newConversation.IsPublic {
		// Add the Logged In User by default
		_, err := c.conversations.AddUserToConversation(context.Background(), &pb.ConversationHasParticipant{
			ConversationId: newConversation.Id,
			ParticipantId:  user.Id,
		})
		if err != nil {
			log.Println(err.Error())
			return stable.NewPostConversationBadRequest().WithPayload(&models.Error{
				Message: fmt.Sprintf("Unable to add %s %s to conversation.", user.FirstName, user.LastName),
			})
		}

		// Now loop through and add the user users
		if len(in.Conversation.Participants) > 0 {
			for _, participant := range in.Conversation.Participants {
				_, err := c.conversations.AddUserToConversation(context.Background(), &pb.ConversationHasParticipant{
					ConversationId: newConversation.Id,
					ParticipantId:  participant.ID.String(),
				})
				if err != nil {
					log.Println(err.Error())
					return stable.NewPostConversationBadRequest().WithPayload(&models.Error{
						Message: fmt.Sprintf("Unable to add %s %s to conversation.", participant.FirstName, participant.LastName),
					})
				}
			}
		}
	}

	// Convert the Conversation back to Swagger
	rtnConversation := &models.Conversation{}
	if err := pbjson.FromProto(newConversation, rtnConversation); err != nil {
		log.Println(err.Error())
		return stable.NewPostConversationBadRequest().WithPayload(&models.Error{
			Message: "Unable to convert conversation information.",
		})
	}

	// Return the Conversation
	return stable.NewPostConversationCreated().WithPayload(rtnConversation)
}
