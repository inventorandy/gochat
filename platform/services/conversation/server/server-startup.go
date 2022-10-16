package server

import (
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/conversation/controller"
	"gochat/platform/services/conversation/interfaces"
	"sync"
)

// ChatEventConnection struct
type ChatEventConnection struct {
	stream pb.ConversationService_ChatStreamServer
	active bool
	err    chan error
}

// ConversationService struct
type ConversationService struct {
	c                    interfaces.Controller
	ChatEventConnections []*ChatEventConnection
	chatMtx              sync.Mutex
	pb.ConversationServiceServer
}

// NewConversationService creates a new instance of the Conversation gRPC Server
func NewConversationService() (*ConversationService, error) {
	// Start up the Controller
	controller, err := controller.NewConversationController()
	if err != nil {
		return nil, err
	}

	// Return the new gRPC Server
	return &ConversationService{
		c: controller,
	}, nil
}
