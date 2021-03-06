package main

import (
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/gorilla/websocket"
)

// Client wraps a single connection to the websocket server
type Client struct {
	conn *websocket.Conn
	jwt  string
	send chan Event
}

// WebSocketClientHandler manages connections between clients and internal gRPC streams
type WebSocketClientHandler struct {
	accounts       pb.AccountServiceClient
	conversations  pb.ConversationServiceClient
	authMiddleware func(Event, string) bool
	clients        map[*Client]bool
	register       chan *Client
	unregister     chan *Client
	broadcast      chan Event
}

// Event is the interface for a websocket event
type Event struct {
	Type      string      `json:"type"`
	EventType string      `json:"event_type"`
	Data      interface{} `json:"data"`
}

func (s *WebSocketClientHandler) run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				log.Println("Unregistering client")
				close(client.send)
				delete(s.clients, client)
			}
		case message := <-s.broadcast:
			for client := range s.clients {
				if s.authMiddleware(message, client.jwt) {
					select {
					case client.send <- message:
					default:
						log.Println("Unregistering client")
						close(client.send)
						delete(s.clients, client)
					}
				}
			}
		}
	}
}

// run handler sends data to a client
func (c *Client) run() {
	for {
		message := <-c.send
		c.conn.WriteJSON(message)
	}
}
