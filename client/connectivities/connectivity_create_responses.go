// Code generated by go-swagger; DO NOT EDIT.

package connectivities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// ConnectivityCreateReader is a Reader for the ConnectivityCreate structure.
type ConnectivityCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectivityCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectivityCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectivityCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectivityCreateOK creates a ConnectivityCreateOK with default headers values
func NewConnectivityCreateOK() *ConnectivityCreateOK {
	return &ConnectivityCreateOK{}
}

/*
ConnectivityCreateOK describes a response with status code 200, with default header values.

New connectivity was created successfully
*/
type ConnectivityCreateOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this connectivity create o k response has a 2xx status code
func (o *ConnectivityCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this connectivity create o k response has a 3xx status code
func (o *ConnectivityCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connectivity create o k response has a 4xx status code
func (o *ConnectivityCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this connectivity create o k response has a 5xx status code
func (o *ConnectivityCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this connectivity create o k response a status code equal to that given
func (o *ConnectivityCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the connectivity create o k response
func (o *ConnectivityCreateOK) Code() int {
	return 200
}

func (o *ConnectivityCreateOK) Error() string {
	return fmt.Sprintf("[POST /connectivities][%d] connectivityCreateOK  %+v", 200, o.Payload)
}

func (o *ConnectivityCreateOK) String() string {
	return fmt.Sprintf("[POST /connectivities][%d] connectivityCreateOK  %+v", 200, o.Payload)
}

func (o *ConnectivityCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *ConnectivityCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectivityCreateDefault creates a ConnectivityCreateDefault with default headers values
func NewConnectivityCreateDefault(code int) *ConnectivityCreateDefault {
	return &ConnectivityCreateDefault{
		_statusCode: code,
	}
}

/*
ConnectivityCreateDefault describes a response with status code -1, with default header values.

Error creating new connectivity
*/
type ConnectivityCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this connectivity create default response has a 2xx status code
func (o *ConnectivityCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this connectivity create default response has a 3xx status code
func (o *ConnectivityCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this connectivity create default response has a 4xx status code
func (o *ConnectivityCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this connectivity create default response has a 5xx status code
func (o *ConnectivityCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this connectivity create default response a status code equal to that given
func (o *ConnectivityCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the connectivity create default response
func (o *ConnectivityCreateDefault) Code() int {
	return o._statusCode
}

func (o *ConnectivityCreateDefault) Error() string {
	return fmt.Sprintf("[POST /connectivities][%d] ConnectivityCreate default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectivityCreateDefault) String() string {
	return fmt.Sprintf("[POST /connectivities][%d] ConnectivityCreate default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectivityCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConnectivityCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
