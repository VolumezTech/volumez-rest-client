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

// NetworkGetReader is a Reader for the NetworkGet structure.
type NetworkGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkGetOK creates a NetworkGetOK with default headers values
func NewNetworkGetOK() *NetworkGetOK {
	return &NetworkGetOK{}
}

/*
NetworkGetOK describes a response with status code 200, with default header values.

Properties of a network
*/
type NetworkGetOK struct {
	Payload *models.Network
}

func (o *NetworkGetOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network}][%d] networkGetOK  %+v", 200, o.Payload)
}
func (o *NetworkGetOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *NetworkGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkGetDefault creates a NetworkGetDefault with default headers values
func NewNetworkGetDefault(code int) *NetworkGetDefault {
	return &NetworkGetDefault{
		_statusCode: code,
	}
}

/*
NetworkGetDefault describes a response with status code -1, with default header values.

Error getting properties of a network
*/
type NetworkGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network get default response
func (o *NetworkGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkGetDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network}][%d] NetworkGet default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
