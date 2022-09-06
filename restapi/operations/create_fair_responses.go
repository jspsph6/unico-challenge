// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"unico-challenge/models"
)

// CreateFairOKCode is the HTTP code returned for type CreateFairOK
const CreateFairOKCode int = 200

/*CreateFairOK Successful operation

swagger:response createFairOK
*/
type CreateFairOK struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewCreateFairOK creates CreateFairOK with default headers values
func NewCreateFairOK() *CreateFairOK {

	return &CreateFairOK{}
}

// WithPayload adds the payload to the create fair o k response
func (o *CreateFairOK) WithPayload(payload *models.Response) *CreateFairOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create fair o k response
func (o *CreateFairOK) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFairOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFairBadRequestCode is the HTTP code returned for type CreateFairBadRequest
const CreateFairBadRequestCode int = 400

/*CreateFairBadRequest One or more fields is invalid

swagger:response createFairBadRequest
*/
type CreateFairBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewCreateFairBadRequest creates CreateFairBadRequest with default headers values
func NewCreateFairBadRequest() *CreateFairBadRequest {

	return &CreateFairBadRequest{}
}

// WithPayload adds the payload to the create fair bad request response
func (o *CreateFairBadRequest) WithPayload(payload *models.Response) *CreateFairBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create fair bad request response
func (o *CreateFairBadRequest) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFairBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
