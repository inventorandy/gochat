package controller

import (
	"fmt"
	"gochat/platform/internal/proto/pb"
	"os"

	"google.golang.org/grpc"
)

// HandlerController struct
type HandlerController struct {
	accounts      pb.AccountServiceClient
	conversations pb.ConversationServiceClient
}

// NewHandlerController creates a new instance of the Handler Controller
func NewHandlerController() (*HandlerController, error) {
	// Get the Accounts Service Host
	accountHost, exists := os.LookupEnv("ACCOUNT_SERVICE_HOST")
	if !exists {
		return nil, fmt.Errorf("no ACCOUNT_SERVICE_HOST specified")
	}

	// Get the Accounts Service Port
	accountPort, exists := os.LookupEnv("ACCOUNT_SERVICE_PORT")
	if !exists {
		return nil, fmt.Errorf("no ACCOUNT_SERVICE_PORT specified")
	}

	// Build the Connection String
	accountConnString := fmt.Sprintf("%s:%s", accountHost, accountPort)

	// Connect to the Service
	accountConn, err := grpc.Dial(accountConnString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	// Create the Client
	accountClient := pb.NewAccountServiceClient(accountConn)

	// Get the Conversations Service Host
	conversationHost, exists := os.LookupEnv("CONVERSATION_SERVICE_HOST")
	if !exists {
		return nil, fmt.Errorf("no CONVERSATION_SERVICE_HOST specified")
	}

	// Get the Conversations Service Port
	conversationPort, exists := os.LookupEnv("CONVERSATION_SERVICE_PORT")
	if !exists {
		return nil, fmt.Errorf("no CONVERSATION_SERVICE_PORT specified")
	}

	// Build the Connection String
	conversationConnString := fmt.Sprintf("%s:%s", conversationHost, conversationPort)

	// Connect to the Service
	conversationConn, err := grpc.Dial(conversationConnString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	// Create the Client
	conversationClient := pb.NewConversationServiceClient(conversationConn)

	// Return the Controller
	return &HandlerController{
		accounts:      accountClient,
		conversations: conversationClient,
	}, nil
}
