// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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

// IsSuccess returns true when this network get o k response has a 2xx status code
func (o *NetworkGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network get o k response has a 3xx status code
func (o *NetworkGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network get o k response has a 4xx status code
func (o *NetworkGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network get o k response has a 5xx status code
func (o *NetworkGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network get o k response a status code equal to that given
func (o *NetworkGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network get o k response
func (o *NetworkGetOK) Code() int {
	return 200
}

func (o *NetworkGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /networks/{network}][%d] networkGetOK %s", 200, payload)
}

func (o *NetworkGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /networks/{network}][%d] networkGetOK %s", 200, payload)
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

// IsSuccess returns true when this network get default response has a 2xx status code
func (o *NetworkGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network get default response has a 3xx status code
func (o *NetworkGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network get default response has a 4xx status code
func (o *NetworkGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network get default response has a 5xx status code
func (o *NetworkGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network get default response a status code equal to that given
func (o *NetworkGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network get default response
func (o *NetworkGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /networks/{network}][%d] NetworkGet default %s", o._statusCode, payload)
}

func (o *NetworkGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /networks/{network}][%d] NetworkGet default %s", o._statusCode, payload)
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
