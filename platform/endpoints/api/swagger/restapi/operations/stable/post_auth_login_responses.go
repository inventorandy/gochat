// Code generated by go-swagger; DO NOT EDIT.

package stable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochat/platform/endpoints/api/swagger/models"
)

// PostAuthLoginOKCode is the HTTP code returned for type PostAuthLoginOK
const PostAuthLoginOKCode int = 200

/*PostAuthLoginOK User was successfully logged in.

swagger:response postAuthLoginOK
*/
type PostAuthLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewPostAuthLoginOK creates PostAuthLoginOK with default headers values
func NewPostAuthLoginOK() *PostAuthLoginOK {

	return &PostAuthLoginOK{}
}

// WithPayload adds the payload to the post auth login o k response
func (o *PostAuthLoginOK) WithPayload(payload *models.LoginResponse) *PostAuthLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login o k response
func (o *PostAuthLoginOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthLoginUnauthorizedCode is the HTTP code returned for type PostAuthLoginUnauthorized
const PostAuthLoginUnauthorizedCode int = 401

/*PostAuthLoginUnauthorized Unauthorized

swagger:response postAuthLoginUnauthorized
*/
type PostAuthLoginUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostAuthLoginUnauthorized creates PostAuthLoginUnauthorized with default headers values
func NewPostAuthLoginUnauthorized() *PostAuthLoginUnauthorized {

	return &PostAuthLoginUnauthorized{}
}

// WithPayload adds the payload to the post auth login unauthorized response
func (o *PostAuthLoginUnauthorized) WithPayload(payload interface{}) *PostAuthLoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login unauthorized response
func (o *PostAuthLoginUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostAuthLoginInternalServerErrorCode is the HTTP code returned for type PostAuthLoginInternalServerError
const PostAuthLoginInternalServerErrorCode int = 500

/*PostAuthLoginInternalServerError Unknown Error

swagger:response postAuthLoginInternalServerError
*/
type PostAuthLoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostAuthLoginInternalServerError creates PostAuthLoginInternalServerError with default headers values
func NewPostAuthLoginInternalServerError() *PostAuthLoginInternalServerError {

	return &PostAuthLoginInternalServerError{}
}

// WithPayload adds the payload to the post auth login internal server error response
func (o *PostAuthLoginInternalServerError) WithPayload(payload interface{}) *PostAuthLoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login internal server error response
func (o *PostAuthLoginInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
