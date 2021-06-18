package server

import (
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

// ChatStream gRPC Stream Method
func (s *ConversationService) ChatStream(in *empty.Empty, srv pb.ConversationService_ChatStreamServer) error {
	// Create the Chat Event Connection
	conn := &ChatEventConnection{
		stream: srv,
		active: true,
		err:    make(chan error),
	}

	// Add the Connection to the List
	s.chatMtx.Lock()
	defer s.chatMtx.Unlock()
	s.ChatEventConnections = append(s.ChatEventConnections, conn)

	// Return any error sent
	return <-conn.err
}

// sendChatEvent publishes a chat event to the stream
func (s *ConversationService) sendChatEvent(in *pb.ChatStreamUpdate) {
	s.chatMtx.Lock()
	defer s.chatMtx.Unlock()

	// Send data and close any inactive connections
	for _, conn := range s.ChatEventConnections {
		go func(in *pb.ChatStreamUpdate, conn *ChatEventConnection) {
			if conn.active {
				if err := conn.stream.Send(in); err != nil {
					conn.active = false
					conn.err <- err
				}
			}
		}(in, conn)
	}
}
