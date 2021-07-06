package main

import (
	"fmt"
	"gochat/platform/internal/proto/pb"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

// Upgrade the Connection to a Websocket Connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// main Entrypoint
func main() {
	// Get the Port
	wsPort, exists := os.LookupEnv("WS_SERVICE_PORT")
	if !exists {
		// Shut down
		log.Fatalln("No WS_SERVICE_PORT environment variable specified")
	}

	// Log a message
	log.Println("Starting websocket service on port", wsPort)

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

	// Create the Client
	accountClient := pb.NewAccountServiceClient(accountConn)

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

	// Create the Client
	conversationClient := pb.NewConversationServiceClient(conversationConn)

	// Set up the Conversations Websocket Stream
	conversationSocketHandler := &WebSocketClientHandler{
		accounts:      accountClient,
		conversations: conversationClient,
		clients:       make(map[*Client]bool),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		broadcast:     make(chan interface{}),
	}
	go conversationSocketHandler.run()
	http.HandleFunc("/conversations", conversationSocketHandler.ConversationHandler)
	go conversationSocketHandler.listenForConversationEvents()

	// Set up the Server and Serve
	address := fmt.Sprintf(":%s", wsPort)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalln(err.Error())
	}
}
