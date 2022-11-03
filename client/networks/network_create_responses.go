// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"bitbucket.org/volumez/volumez/test/restapi/models"
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

/* NetworkCreateOK describes a response with status code 200, with default header values.

New network was created successfully
*/
type NetworkCreateOK struct {
	Payload *models.RegularResponse
}

func (o *NetworkCreateOK) Error() string {
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

/* NetworkCreateDefault describes a response with status code -1, with default header values.

Error creating new network
*/
type NetworkCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network create default response
func (o *NetworkCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkCreateDefault) Error() string {
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
