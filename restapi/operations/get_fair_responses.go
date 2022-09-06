// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"unico-challenge/models"
)

// GetFairOKCode is the HTTP code returned for type GetFairOK
const GetFairOKCode int = 200

/*GetFairOK Successful operation

swagger:response getFairOK
*/
type GetFairOK struct {

	/*
	  In: Body
	*/
	Payload *models.Fair `json:"body,omitempty"`
}

// NewGetFairOK creates GetFairOK with default headers values
func NewGetFairOK() *GetFairOK {

	return &GetFairOK{}
}

// WithPayload adds the payload to the get fair o k response
func (o *GetFairOK) WithPayload(payload *models.Fair) *GetFairOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fair o k response
func (o *GetFairOK) SetPayload(payload *models.Fair) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFairOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFairBadRequestCode is the HTTP code returned for type GetFairBadRequest
const GetFairBadRequestCode int = 400

/*GetFairBadRequest One or more fields is invalid

swagger:response getFairBadRequest
*/
type GetFairBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetFairBadRequest creates GetFairBadRequest with default headers values
func NewGetFairBadRequest() *GetFairBadRequest {

	return &GetFairBadRequest{}
}

// WithPayload adds the payload to the get fair bad request response
func (o *GetFairBadRequest) WithPayload(payload *models.Response) *GetFairBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fair bad request response
func (o *GetFairBadRequest) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFairBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}