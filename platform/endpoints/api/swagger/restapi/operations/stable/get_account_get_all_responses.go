// Code generated by go-swagger; DO NOT EDIT.

package stable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochat/platform/endpoints/api/swagger/models"
)

// GetAccountGetAllOKCode is the HTTP code returned for type GetAccountGetAllOK
const GetAccountGetAllOKCode int = 200

/*GetAccountGetAllOK User list was returned.

swagger:response getAccountGetAllOK
*/
type GetAccountGetAllOK struct {

	/*
	  In: Body
	*/
	Payload models.UserList `json:"body,omitempty"`
}

// NewGetAccountGetAllOK creates GetAccountGetAllOK with default headers values
func NewGetAccountGetAllOK() *GetAccountGetAllOK {

	return &GetAccountGetAllOK{}
}

// WithPayload adds the payload to the get account get all o k response
func (o *GetAccountGetAllOK) WithPayload(payload models.UserList) *GetAccountGetAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account get all o k response
func (o *GetAccountGetAllOK) SetPayload(payload models.UserList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountGetAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.UserList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountGetAllUnauthorizedCode is the HTTP code returned for type GetAccountGetAllUnauthorized
const GetAccountGetAllUnauthorizedCode int = 401

/*GetAccountGetAllUnauthorized Unauthorized

swagger:response getAccountGetAllUnauthorized
*/
type GetAccountGetAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAccountGetAllUnauthorized creates GetAccountGetAllUnauthorized with default headers values
func NewGetAccountGetAllUnauthorized() *GetAccountGetAllUnauthorized {

	return &GetAccountGetAllUnauthorized{}
}

// WithPayload adds the payload to the get account get all unauthorized response
func (o *GetAccountGetAllUnauthorized) WithPayload(payload interface{}) *GetAccountGetAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account get all unauthorized response
func (o *GetAccountGetAllUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountGetAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountGetAllInternalServerErrorCode is the HTTP code returned for type GetAccountGetAllInternalServerError
const GetAccountGetAllInternalServerErrorCode int = 500

/*GetAccountGetAllInternalServerError Unknown Error

swagger:response getAccountGetAllInternalServerError
*/
type GetAccountGetAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAccountGetAllInternalServerError creates GetAccountGetAllInternalServerError with default headers values
func NewGetAccountGetAllInternalServerError() *GetAccountGetAllInternalServerError {

	return &GetAccountGetAllInternalServerError{}
}

// WithPayload adds the payload to the get account get all internal server error response
func (o *GetAccountGetAllInternalServerError) WithPayload(payload interface{}) *GetAccountGetAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account get all internal server error response
func (o *GetAccountGetAllInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountGetAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}