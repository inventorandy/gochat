package main

import (
	"context"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"
	"log"
	"reflect"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func (s *WebSocketClientHandler) listenForConversationEvents() {
	// Create the stream for conversation events
	stream, err := s.conversations.ChatStream(context.Background(), &empty.Empty{}, grpc.WaitForReady(true))
	if err != nil {
		log.Println(err.Error())
		go s.listenForConversationEvents()
		return
	}

	// Print a quick log
	log.Println("Conversation stream established...")

	// Main Listener Loop
	for {
		// Get the Event
		event, err := stream.Recv()
		if err != nil {
			log.Println(err.Error())
			go s.listenForConversationEvents()
			return
		}

		// Print a quick log
		log.Printf("Conversation Event: %+v\r\n", event)

		// Determine the Event Type
		switch reflect.TypeOf(event.UpdateType) {
		case reflect.TypeOf(&pb.ChatStreamUpdate_MessageInfo{}):
			// Get the Message Info
			msgInfo := event.GetMessageInfo()

			// Convert the Message to JSON
			msgJSON := &models.Message{}
			if err := pbjson.FromProto(msgInfo.Message, msgJSON); err != nil {
				log.Println(err.Error())
				continue
			}

			// Broadcast the Event
			s.broadcast <- Event{
				Type: msgInfo.UpdateType.String(),
				Data: msgJSON,
			}
		default:
			log.Println("ignoring event...")
			continue
		}
	}
}
