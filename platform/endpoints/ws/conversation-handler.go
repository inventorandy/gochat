package main

import (
	"log"
	"net/http"
)

// ConversationHandler deals with websocket connections to [/conversation]
func (s *WebSocketClientHandler) ConversationHandler(w http.ResponseWriter, r *http.Request) {
	// Get the User
	authToken, err := r.Cookie("token")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Upgrade the HTTP connection to WS
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket error: ", err.Error())
		return
	}

	// Create the Websocket Client
	wsClient := &Client{
		conn: conn,
		jwt:  authToken.Value,
		send: make(chan Event, 256),
	}

	// Register the new Client
	s.register <- wsClient

	// Run the Listener
	go wsClient.run()
}
