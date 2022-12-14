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

// ConnectivityModifyReader is a Reader for the ConnectivityModify structure.
type ConnectivityModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectivityModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectivityModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectivityModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectivityModifyOK creates a ConnectivityModifyOK with default headers values
func NewConnectivityModifyOK() *ConnectivityModifyOK {
	return &ConnectivityModifyOK{}
}

/* ConnectivityModifyOK describes a response with status code 200, with default header values.

A connectivity was updated successfully
*/
type ConnectivityModifyOK struct {
	Payload *models.RegularResponse
}

func (o *ConnectivityModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /connectivities/{connectivity}][%d] connectivityModifyOK  %+v", 200, o.Payload)
}
func (o *ConnectivityModifyOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *ConnectivityModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectivityModifyDefault creates a ConnectivityModifyDefault with default headers values
func NewConnectivityModifyDefault(code int) *ConnectivityModifyDefault {
	return &ConnectivityModifyDefault{
		_statusCode: code,
	}
}

/* ConnectivityModifyDefault describes a response with status code -1, with default header values.

Error updating a connectivity
*/
type ConnectivityModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the connectivity modify default response
func (o *ConnectivityModifyDefault) Code() int {
	return o._statusCode
}

func (o *ConnectivityModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /connectivities/{connectivity}][%d] ConnectivityModify default  %+v", o._statusCode, o.Payload)
}
func (o *ConnectivityModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConnectivityModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
