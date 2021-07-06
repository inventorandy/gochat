package main

import (
	"gochat/platform/internal/proto/pb"
	"log"

	"github.com/gorilla/websocket"
)

// Client wraps a single connection to the websocket server
type Client struct {
	conn *websocket.Conn
	send chan interface{}
}

// WebSocketClientHandler manages connections between clients and internal gRPC streams
type WebSocketClientHandler struct {
	accounts      pb.AccountServiceClient
	conversations pb.ConversationServiceClient
	clients       map[*Client]bool
	register      chan *Client
	unregister    chan *Client
	broadcast     chan interface{}
}

// Event is the interface for a websocket event
type Event struct {
	Type string
	Data interface{}
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
