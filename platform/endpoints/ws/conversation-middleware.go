package main

import (
	"context"
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

// AuthConversation Middleware
func AuthConversation(event Event, jwt string) bool {
	// Get the Conversations Service Host
	conversationHost, exists := os.LookupEnv("CONVERSATION_SERVICE_HOST")
	if !exists {
		log.Fatalln("no CONVERSATION_SERVICE_HOST specified")
	}

	// Get the Conversations Service Port
	conversationPort, exists := os.LookupEnv("CONVERSATION_SERVICE_PORT")
	if !exists {
		log.Fatalln("no CONVERSATION_SERVICE_PORT specified")
	}

	// Build the Connection String
	conversationConnString := fmt.Sprintf("%s:%s", conversationHost, conversationPort)

	// Connect to the Service
	conversationConn, err := grpc.Dial(conversationConnString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conversationConn.Close()

	// Create the Client
	conversationClient := pb.NewConversationServiceClient(conversationConn)

	// Get the Accounts Service Host
	accountHost, exists := os.LookupEnv("ACCOUNT_SERVICE_HOST")
	if !exists {
		log.Fatalln("no ACCOUNT_SERVICE_HOST specified")
	}

	// Get the Accounts Service Port
	accountPort, exists := os.LookupEnv("ACCOUNT_SERVICE_PORT")
	if !exists {
		log.Fatalln("no ACCOUNT_SERVICE_PORT specified")
	}

	// Build the Connection String
	accountConnString := fmt.Sprintf("%s:%s", accountHost, accountPort)

	// Connect to the Service
	accountConn, err := grpc.Dial(accountConnString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer accountConn.Close()

	// Create the Client
	accountClient := pb.NewAccountServiceClient(accountConn)

	// Get the User
	user, err := accountClient.AuthenticateJWT(context.Background(), &wrappers.StringValue{Value: jwt})
	if err != nil {
		log.Println(err.Error())
		return false
	}

	// Determine the Message Type
	switch event.Type {
	case "MESSAGE":
		// Convert to Proto
		update := &pb.Message{}
		if err := pbjson.ToProto(event.Data, update); err != nil {
			log.Println(err.Error())
			return false
		}

		// Check whether the user has access to the conversation
		userHasAccess, err := conversationClient.UserHasAccessToConversation(context.Background(), &pb.UserAccessQuery{
			ConversationId: update.ConversationId,
			UserId:         user.Id,
		})
		if err != nil {
			log.Println(err.Error())
			return false
		}

		// Return the Value
		return userHasAccess.Value
	case "CONVERSATION":
		// Convert to Proto
		update := &pb.Conversation{}
		if err := pbjson.ToProto(event.Data, update); err != nil {
			log.Println(err.Error())
			return false
		}

		// Check whether the user has access to the conversation
		userHasAccess, err := conversationClient.UserHasAccessToConversation(context.Background(), &pb.UserAccessQuery{
			ConversationId: update.Id,
			UserId:         user.Id,
		})
		if err != nil {
			log.Println(err.Error())
			return false
		}

		// Return the Value
		return userHasAccess.Value
	default:
		// Return False by Default
		return false
	}
}
