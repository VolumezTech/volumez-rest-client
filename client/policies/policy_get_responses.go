// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// PolicyGetReader is a Reader for the PolicyGet structure.
type PolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyGetOK creates a PolicyGetOK with default headers values
func NewPolicyGetOK() *PolicyGetOK {
	return &PolicyGetOK{}
}

/*
PolicyGetOK describes a response with status code 200, with default header values.

Properties of a policy
*/
type PolicyGetOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this policy get o k response has a 2xx status code
func (o *PolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy get o k response has a 3xx status code
func (o *PolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy get o k response has a 4xx status code
func (o *PolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy get o k response has a 5xx status code
func (o *PolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy get o k response a status code equal to that given
func (o *PolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the policy get o k response
func (o *PolicyGetOK) Code() int {
	return 200
}

func (o *PolicyGetOK) Error() string {
	return fmt.Sprintf("[GET /policies/{policy}][%d] policyGetOK  %+v", 200, o.Payload)
}

func (o *PolicyGetOK) String() string {
	return fmt.Sprintf("[GET /policies/{policy}][%d] policyGetOK  %+v", 200, o.Payload)
}

func (o *PolicyGetOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *PolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyGetDefault creates a PolicyGetDefault with default headers values
func NewPolicyGetDefault(code int) *PolicyGetDefault {
	return &PolicyGetDefault{
		_statusCode: code,
	}
}

/*
PolicyGetDefault describes a response with status code -1, with default header values.

Error getting properties of a policy
*/
type PolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this policy get default response has a 2xx status code
func (o *PolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy get default response has a 3xx status code
func (o *PolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy get default response has a 4xx status code
func (o *PolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy get default response has a 5xx status code
func (o *PolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy get default response a status code equal to that given
func (o *PolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the policy get default response
func (o *PolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *PolicyGetDefault) Error() string {
	return fmt.Sprintf("[GET /policies/{policy}][%d] PolicyGet default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyGetDefault) String() string {
	return fmt.Sprintf("[GET /policies/{policy}][%d] PolicyGet default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
