// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montenegrodr/brcovid19api/models"
)

// GetCovid19ReportDataOKCode is the HTTP code returned for type GetCovid19ReportDataOK
const GetCovid19ReportDataOKCode int = 200

/*GetCovid19ReportDataOK Up to date data about covid19 in Brazil

swagger:response getCovid19ReportDataOK
*/
type GetCovid19ReportDataOK struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetCovid19ReportDataOK creates GetCovid19ReportDataOK with default headers values
func NewGetCovid19ReportDataOK() *GetCovid19ReportDataOK {

	return &GetCovid19ReportDataOK{}
}

// WithPayload adds the payload to the get covid19 report data o k response
func (o *GetCovid19ReportDataOK) WithPayload(payload *models.Response) *GetCovid19ReportDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get covid19 report data o k response
func (o *GetCovid19ReportDataOK) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCovid19ReportDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetCovid19ReportDataDefault unexpected error

swagger:response getCovid19ReportDataDefault
*/
type GetCovid19ReportDataDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetCovid19ReportDataDefault creates GetCovid19ReportDataDefault with default headers values
func NewGetCovid19ReportDataDefault(code int) *GetCovid19ReportDataDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCovid19ReportDataDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get covid19 report data default response
func (o *GetCovid19ReportDataDefault) WithStatusCode(code int) *GetCovid19ReportDataDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get covid19 report data default response
func (o *GetCovid19ReportDataDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get covid19 report data default response
func (o *GetCovid19ReportDataDefault) WithPayload(payload *models.ErrorModel) *GetCovid19ReportDataDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get covid19 report data default response
func (o *GetCovid19ReportDataDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCovid19ReportDataDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}