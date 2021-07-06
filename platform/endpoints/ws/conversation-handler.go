package main

import (
	"log"
	"net/http"
)

// ConversationHandler deals with websocket connections to [/conversation]
func (s *WebSocketClientHandler) ConversationHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WS
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket error: ", err.Error())
		return
	}

	// Create the Websocket Client
	wsClient := &Client{
		conn: conn,
		send: make(chan interface{}, 256),
	}

	// Register the new Client
	s.register <- wsClient

	// Run the Listener
	go wsClient.run()
}
