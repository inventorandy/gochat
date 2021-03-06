// Code generated by go-swagger; DO NOT EDIT.

package stable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochat/platform/endpoints/api/swagger/models"
)

// PostConversationCreatedCode is the HTTP code returned for type PostConversationCreated
const PostConversationCreatedCode int = 201

/*PostConversationCreated Conversation was created.

swagger:response postConversationCreated
*/
type PostConversationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Conversation `json:"body,omitempty"`
}

// NewPostConversationCreated creates PostConversationCreated with default headers values
func NewPostConversationCreated() *PostConversationCreated {

	return &PostConversationCreated{}
}

// WithPayload adds the payload to the post conversation created response
func (o *PostConversationCreated) WithPayload(payload *models.Conversation) *PostConversationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post conversation created response
func (o *PostConversationCreated) SetPayload(payload *models.Conversation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConversationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConversationBadRequestCode is the HTTP code returned for type PostConversationBadRequest
const PostConversationBadRequestCode int = 400

/*PostConversationBadRequest Bad Request

swagger:response postConversationBadRequest
*/
type PostConversationBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostConversationBadRequest creates PostConversationBadRequest with default headers values
func NewPostConversationBadRequest() *PostConversationBadRequest {

	return &PostConversationBadRequest{}
}

// WithPayload adds the payload to the post conversation bad request response
func (o *PostConversationBadRequest) WithPayload(payload interface{}) *PostConversationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post conversation bad request response
func (o *PostConversationBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConversationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostConversationUnauthorizedCode is the HTTP code returned for type PostConversationUnauthorized
const PostConversationUnauthorizedCode int = 401

/*PostConversationUnauthorized Unauthorized

swagger:response postConversationUnauthorized
*/
type PostConversationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostConversationUnauthorized creates PostConversationUnauthorized with default headers values
func NewPostConversationUnauthorized() *PostConversationUnauthorized {

	return &PostConversationUnauthorized{}
}

// WithPayload adds the payload to the post conversation unauthorized response
func (o *PostConversationUnauthorized) WithPayload(payload interface{}) *PostConversationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post conversation unauthorized response
func (o *PostConversationUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConversationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostConversationInternalServerErrorCode is the HTTP code returned for type PostConversationInternalServerError
const PostConversationInternalServerErrorCode int = 500

/*PostConversationInternalServerError Unknown Error

swagger:response postConversationInternalServerError
*/
type PostConversationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostConversationInternalServerError creates PostConversationInternalServerError with default headers values
func NewPostConversationInternalServerError() *PostConversationInternalServerError {

	return &PostConversationInternalServerError{}
}

// WithPayload adds the payload to the post conversation internal server error response
func (o *PostConversationInternalServerError) WithPayload(payload interface{}) *PostConversationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post conversation internal server error response
func (o *PostConversationInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConversationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
