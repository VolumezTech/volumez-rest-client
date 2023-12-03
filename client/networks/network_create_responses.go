// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NetworkCreateReader is a Reader for the NetworkCreate structure.
type NetworkCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkCreateOK creates a NetworkCreateOK with default headers values
func NewNetworkCreateOK() *NetworkCreateOK {
	return &NetworkCreateOK{}
}

/*
NetworkCreateOK describes a response with status code 200, with default header values.

New network was created successfully
*/
type NetworkCreateOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this network create o k response has a 2xx status code
func (o *NetworkCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network create o k response has a 3xx status code
func (o *NetworkCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create o k response has a 4xx status code
func (o *NetworkCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create o k response has a 5xx status code
func (o *NetworkCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network create o k response a status code equal to that given
func (o *NetworkCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network create o k response
func (o *NetworkCreateOK) Code() int {
	return 200
}

func (o *NetworkCreateOK) Error() string {
	return fmt.Sprintf("[POST /networks][%d] networkCreateOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateOK) String() string {
	return fmt.Sprintf("[POST /networks][%d] networkCreateOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *NetworkCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateDefault creates a NetworkCreateDefault with default headers values
func NewNetworkCreateDefault(code int) *NetworkCreateDefault {
	return &NetworkCreateDefault{
		_statusCode: code,
	}
}

/*
NetworkCreateDefault describes a response with status code -1, with default header values.

Error creating new network
*/
type NetworkCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network create default response has a 2xx status code
func (o *NetworkCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network create default response has a 3xx status code
func (o *NetworkCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network create default response has a 4xx status code
func (o *NetworkCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network create default response has a 5xx status code
func (o *NetworkCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network create default response a status code equal to that given
func (o *NetworkCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network create default response
func (o *NetworkCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkCreateDefault) Error() string {
	return fmt.Sprintf("[POST /networks][%d] NetworkCreate default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkCreateDefault) String() string {
	return fmt.Sprintf("[POST /networks][%d] NetworkCreate default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
