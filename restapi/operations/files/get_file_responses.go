package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*GetFileOK get file o k

swagger:response getFileOK
*/
type GetFileOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetFileOK creates GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

// WithPayload adds the payload to the get file o k response
func (o *GetFileOK) WithPayload(payload io.ReadCloser) *GetFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get file o k response
func (o *GetFileOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetFileUnauthorized get file unauthorized

swagger:response getFileUnauthorized
*/
type GetFileUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewGetFileUnauthorized creates GetFileUnauthorized with default headers values
func NewGetFileUnauthorized() *GetFileUnauthorized {
	return &GetFileUnauthorized{}
}

// WithPayload adds the payload to the get file unauthorized response
func (o *GetFileUnauthorized) WithPayload(payload *models.Result) *GetFileUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get file unauthorized response
func (o *GetFileUnauthorized) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFileUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFileNotFound get file not found

swagger:response getFileNotFound
*/
type GetFileNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewGetFileNotFound creates GetFileNotFound with default headers values
func NewGetFileNotFound() *GetFileNotFound {
	return &GetFileNotFound{}
}

// WithPayload adds the payload to the get file not found response
func (o *GetFileNotFound) WithPayload(payload *models.Result) *GetFileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get file not found response
func (o *GetFileNotFound) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFileInternalServerError get file internal server error

swagger:response getFileInternalServerError
*/
type GetFileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewGetFileInternalServerError creates GetFileInternalServerError with default headers values
func NewGetFileInternalServerError() *GetFileInternalServerError {
	return &GetFileInternalServerError{}
}

// WithPayload adds the payload to the get file internal server error response
func (o *GetFileInternalServerError) WithPayload(payload *models.Result) *GetFileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get file internal server error response
func (o *GetFileInternalServerError) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}