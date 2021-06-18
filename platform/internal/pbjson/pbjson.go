package pbjson

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// M global marshaler
func M() protojson.MarshalOptions {
	return protojson.MarshalOptions{
		AllowPartial:  true,
		UseProtoNames: true,
	}
}

// U global unmarshaler
func U() protojson.UnmarshalOptions {
	return protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
}

// FromProto converts a Proto Object to the specified output object
func FromProto(in proto.Message, out interface{}) error {
	// Marshal the Proto Object to a JSON String
	b, err := M().Marshal(in)
	if err != nil {
		return err
	}

	// Unmarshal the JSON into the Output
	if err := json.Unmarshal(b, out); err != nil {
		return err
	}

	// Return Nil by Default
	return nil
}

// ToProto converts an input object to a Proto Object
func ToProto(in interface{}, out proto.Message) error {
	// Marshal the Input Object to a JSON String
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	// Unmarshal to Proto
	if err := U().Unmarshal(b, out); err != nil {
		return err
	}

	// Return Nil by Default
	return nil
}
