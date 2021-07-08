package controller

import (
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// Controller interface
type Controller interface {
	// Conversation Methods
	CreateConversation(in *pb.Conversation) (*pb.Conversation, error)
	UpdateConversation(in *pb.Conversation) (*pb.Conversation, error)
	GetConversationByID(in *wrappers.StringValue) (*pb.Conversation, error)
	AddUserToConversation(in *pb.ConversationHasParticipant) (*pb.Conversation, error)
	GetPublicConversations(in *empty.Empty) (*pb.ConversationList, error)
	GetPrivateConversationsForUser(in *wrappers.StringValue) (*pb.ConversationList, error)
	UserHasAccessToConversation(in *pb.UserAccessQuery) (*wrappers.BoolValue, error)

	// Message Methods
	CreateMessage(in *pb.Message) (*pb.Message, error)
}
