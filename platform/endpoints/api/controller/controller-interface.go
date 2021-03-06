package controller

import (
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/proto/pb"

	"github.com/go-openapi/runtime/middleware"
)

// Controller interface
type Controller interface {
	// Authentication
	AuthLoginPost(in stable.PostAuthLoginParams) middleware.Responder
	AuthenticateJWT(jwt string) (*pb.User, error)

	// Account
	AccountPost(in stable.PostAccountParams) middleware.Responder
	AccountPut(user *pb.User, in stable.PutAccountParams) middleware.Responder
	AccountGet(user *pb.User, in stable.GetAccountParams) middleware.Responder
	AccountGetAll(in stable.GetAccountParams) middleware.Responder

	// Conversation
	ConversationPost(user *pb.User, in stable.PostConversationParams) middleware.Responder
	ConversationPut(in stable.PutConversationParams) middleware.Responder
	ConversationGet(in stable.GetConversationParams) middleware.Responder

	// Message
	MessagePost(user *pb.User, in stable.PostMessageParams) middleware.Responder
}
