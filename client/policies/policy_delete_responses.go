// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"bitbucket.org/volumez/volumez/test/restapi/models"
)

// PolicyDeleteReader is a Reader for the PolicyDelete structure.
type PolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyDeleteOK creates a PolicyDeleteOK with default headers values
func NewPolicyDeleteOK() *PolicyDeleteOK {
	return &PolicyDeleteOK{}
}

/* PolicyDeleteOK describes a response with status code 200, with default header values.

A policy was deleted successfully
*/
type PolicyDeleteOK struct {
	Payload *models.RegularResponse
}

func (o *PolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /policies/{policy}][%d] policyDeleteOK  %+v", 200, o.Payload)
}
func (o *PolicyDeleteOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *PolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyDeleteDefault creates a PolicyDeleteDefault with default headers values
func NewPolicyDeleteDefault(code int) *PolicyDeleteDefault {
	return &PolicyDeleteDefault{
		_statusCode: code,
	}
}

/* PolicyDeleteDefault describes a response with status code -1, with default header values.

Error deleting a policy
*/
type PolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the policy delete default response
func (o *PolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /policies/{policy}][%d] PolicyDelete default  %+v", o._statusCode, o.Payload)
}
func (o *PolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
